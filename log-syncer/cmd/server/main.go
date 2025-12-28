// Package main 日志同步服务入口
package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"zxm_ai_admin/log-syncer/internal/archiver"
	"zxm_ai_admin/log-syncer/internal/config"
	applogger "zxm_ai_admin/log-syncer/internal/logger"
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
		slog.Default().Error("加载配置失败", "error", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()

	// 初始化日志
	logLevel := parseLogLevel(cfg.Log.Level)
	logDir := cfg.Log.Dir
	if err := applogger.System.Init(logDir, logLevel); err != nil {
		slog.Default().Error("系统日志初始化失败", "error", err)
		os.Exit(1)
	}

	applogger.Info("日志同步服务启动",
		"log_dir", cfg.Proxy.LogDir,
		"archive_dir", cfg.Archive.Dir,
		"retention_days", cfg.Archive.RetentionDays,
		"log_service_url", cfg.Server.LogServiceURL,
	)

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
	applogger.Info("启动时执行同步任务")
	runSyncTask(logScanner, logParser, upldr, arch)

	// 等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	applogger.Info("正在关闭服务...")
	sched.Stop()
	applogger.Info("服务已关闭")
}

// runSyncTask 执行同步任务
func runSyncTask(
	logScanner *scanner.Scanner,
	logParser *parser.Parser,
	upldr *uploader.Uploader,
	arch *archiver.Archiver,
) {
	applogger.Info("===== 开始扫描日志文件 =====")

	cfg := config.GetConfig()
	cutoffTime := getCutoffTime()

	// 扫描日志文件
	files, err := logScanner.Scan(cutoffTime)
	if err != nil {
		applogger.Error("扫描日志文件失败", "error", err)
		return
	}

	if len(files) == 0 {
		applogger.Info("没有需要上传的日志文件")
	} else {
		applogger.Info("发现日志文件", "count", len(files))
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

	applogger.Info("日志分类完成",
		"request_count", len(requestFiles),
		"system_count", len(systemFiles),
	)

	// 处理请求日志
	for _, file := range requestFiles {
		if err := processLogFile(file, logParser, upldr, arch, cfg.Uploader.BatchSize, true); err != nil {
			applogger.Error("处理请求日志失败", "file", file.Name, "error", err)
		}
	}

	// 处理系统日志
	for _, file := range systemFiles {
		if err := processLogFile(file, logParser, upldr, arch, cfg.Uploader.BatchSize, false); err != nil {
			applogger.Error("处理系统日志失败", "file", file.Name, "error", err)
		}
	}

	applogger.Info("===== 扫描完成 =====")

	// 清理过期归档
	if err := arch.CleanExpired(); err != nil {
		applogger.Error("清理过期归档失败", "error", err)
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
	applogger.Info("处理文件", "file", file.Name)

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
		applogger.Info("文件为空或无有效日志", "file", file.Name)
		// 空文件也归档
		return arch.Archive(file.Path)
	}

	applogger.Info("解析日志", "file", file.Name, "entries", len(entries))

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
			applogger.Error("批次上传失败",
				"batch", i+1,
				"total", len(batches),
				"error", result.Error,
			)
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

// parseLogLevel 解析日志级别
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
