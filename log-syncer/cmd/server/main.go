// Package main 日志同步服务入口
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zxm_ai_admin/log-syncer/internal/archiver"
	"zxm_ai_admin/log-syncer/internal/config"
	"zxm_ai_admin/log-syncer/internal/parser"
	"zxm_ai_admin/log-syncer/internal/scheduler"
	"zxm_ai_admin/log-syncer/internal/scanner"
	"zxm_ai_admin/log-syncer/internal/uploader"
)

func main() {
	// 加载配置
	configPath := "configs/config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if err := config.Load(configPath); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	cfg := config.GetConfig()

	log.Println("日志同步服务启动")
	log.Printf("日志目录: %s", cfg.Proxy.LogDir)
	log.Printf("归档目录: %s", cfg.Archive.Dir)
	log.Printf("归档保留: %d 天", cfg.Archive.RetentionDays)
	log.Printf("Log Service: %s", cfg.Server.LogServiceURL)

	// 创建组件
	logScanner := scanner.NewScanner(cfg.Proxy.LogDir)
	logParser := parser.NewParser()
	upldr := uploader.NewUploader(cfg.Server.LogServiceURL, cfg.Server.SystemAuthToken, cfg.Server.Timeout)
	arch := archiver.NewArchiver(cfg.Archive.Dir, cfg.Archive.RetentionDays)

	// 创建调度器
	sched := scheduler.NewScheduler(func() {
		runSyncTask(logScanner, logParser, upldr, arch)
	})

	// 启动调度器
	sched.Start()

	// 启动时立即执行一次同步任务
	log.Println("启动时执行同步任务")
	runSyncTask(logScanner, logParser, upldr, arch)

	// 等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务...")
	sched.Stop()
	log.Println("服务已关闭")
}

// runSyncTask 执行同步任务
func runSyncTask(
	logScanner *scanner.Scanner,
	logParser *parser.Parser,
	upldr *uploader.Uploader,
	arch *archiver.Archiver,
) {
	log.Println("===== 开始扫描日志文件 =====")

	cfg := config.GetConfig()
	cutoffTime := getCutoffTime()

	// 扫描日志文件
	files, err := logScanner.Scan(cutoffTime)
	if err != nil {
		log.Printf("扫描日志文件失败: %v", err)
		return
	}

	if len(files) == 0 {
		log.Println("没有需要上传的日志文件")
	} else {
		log.Printf("发现 %d 个日志文件", len(files))
	}

	// 按类型分类
	var requestFiles []*scanner.LogFile
	var systemFiles []*scanner.LogFile

	for _, file := range files {
		if file.Type == scanner.LogTypeRequest {
			requestFiles = append(requestFiles, file)
		} else {
			systemFiles = append(systemFiles, file)
		}
	}

	log.Printf("请求日志: %d 个, 系统日志: %d 个", len(requestFiles), len(systemFiles))

	// 处理请求日志
	for _, file := range requestFiles {
		if err := processLogFile(file, logParser, upldr, arch, cfg.Uploader.BatchSize, true); err != nil {
			log.Printf("处理请求日志失败: %s, 错误: %v", file.Name, err)
		}
	}

	// 处理系统日志
	for _, file := range systemFiles {
		if err := processLogFile(file, logParser, upldr, arch, cfg.Uploader.BatchSize, false); err != nil {
			log.Printf("处理系统日志失败: %s, 错误: %v", file.Name, err)
		}
	}

	log.Println("===== 扫描完成 =====")

	// 清理过期归档
	if err := arch.CleanExpired(); err != nil {
		log.Printf("清理过期归档失败: %v", err)
	}
}

// processLogFile 处理单个日志文件
func processLogFile(
	file *scanner.LogFile,
	p *parser.Parser,
	upldr *uploader.Uploader,
	arch *archiver.Archiver,
	batchSize int,
	isRequest bool,
) error {
	log.Printf("处理文件: %s", file.Name)

	// 解析日志文件
	logType := parser.LogTypeSystem
	if isRequest {
		logType = parser.LogTypeRequest
	}

	entries, err := p.ParseFile(file.Path, logType)
	if err != nil {
		return err
	}

	if len(entries) == 0 {
		log.Printf("文件为空或无有效日志: %s", file.Name)
		// 空文件也归档
		return arch.Archive(file.Path)
	}

	log.Printf("解析日志: %s (%d 条)", file.Name, len(entries))

	// 分批上传
	batches := p.Batch(entries, batchSize)
	successCount := 0

	for i, batch := range batches {
		var result uploader.UploadResult
		if isRequest {
			// 转换为 []interface{}
			batchEntries := make([]interface{}, len(batch))
			for j, entry := range batch {
				batchEntries[j] = entry
			}
			result = upldr.UploadWithRetry(batchEntries, true)
		} else {
			// 转换系统日志为 []interface{}
			batchEntries := make([]interface{}, len(batch))
			for j, entry := range batch {
				batchEntries[j] = entry
			}
			result = upldr.UploadWithRetry(batchEntries, false)
		}

		if result.Success {
			successCount++
		} else {
			log.Printf("批次 %d/%d 上传失败: %v", i+1, len(batches), result.Error)
			return result.Error // 失败不重试，直接返回
		}
	}

	// 上传成功后归档
	if successCount == len(batches) {
		return arch.Archive(file.Path)
	}

	return nil
}

// getCutoffTime 获取截止时间（当前时间）
func getCutoffTime() time.Time {
	return time.Now()
}
