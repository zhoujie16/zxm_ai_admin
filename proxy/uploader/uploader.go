// Package uploader 日志文件上传器
// 定期扫描日志目录，上传已完成的日志文件到 log-service
package uploader

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"proxy/logger"
)

const (
	// 上传间隔：30 分钟
	uploadInterval = 30 * time.Minute
)

// Uploader 日志上传器
type Uploader struct {
	logDir       string
	serviceURL   string
	apiKey       string
	stopCh       chan struct{}
	httpClient   *http.Client
	uploadingMu  sync.Mutex  // 防止并发上传
}

// New 创建日志上传器
func New(logDir, serviceURL, apiKey string) *Uploader {
	if serviceURL == "" {
		return nil
	}
	return &Uploader{
		logDir:     logDir,
		serviceURL: serviceURL,
		apiKey:     apiKey,
		stopCh:     make(chan struct{}),
		httpClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
}

// Start 启动上传器
func (u *Uploader) Start() {
	if u == nil {
		return
	}

	// 立即执行一次上传
	go u.tryUpload()

	// 定时上传
	go func() {
		ticker := time.NewTicker(uploadInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				u.tryUpload()
			case <-u.stopCh:
				return
			}
		}
	}()

	logger.Info("日志上传器已启动", "interval", uploadInterval)
}

// Stop 停止上传器
func (u *Uploader) Stop() {
	if u != nil && u.stopCh != nil {
		close(u.stopCh)
		logger.Info("日志上传器已停止")
	}
}

// tryUpload 尝试上传（使用互斥锁防止并发）
func (u *Uploader) tryUpload() {
	if u == nil {
		return
	}

	// 尝试获取锁，如果已有上传任务在进行则跳过
	if !u.uploadingMu.TryLock() {
		return
	}
	defer u.uploadingMu.Unlock()

	u.doUpload()
}

// getCurrentHalfHour 获取当前半小时标识
func getCurrentHalfHour() string {
	now := time.Now()
	minute := now.Minute()
	halfHour := now.Hour()*2 + minute/30
	return now.Format("2006-01-02_") + fmt.Sprintf("%02d", halfHour)
}

// doUpload 扫描并上传日志文件（由 tryUpload 调用，已持有锁）
func (u *Uploader) doUpload() {
	if u == nil {
		return
	}

	// 扫描日志文件
	files, err := os.ReadDir(u.logDir)
	if err != nil {
		logger.Error("扫描日志目录失败", "error", err)
		return
	}

	currentHalfHour := getCurrentHalfHour()
	uploadedCount := 0
	failedCount := 0

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		// 检查是否是日志文件
		logType, fileHalfHour, ok := parseLogFilename(filename)
		if !ok {
			continue
		}

		// 跳过当前正在写入的文件
		if fileHalfHour >= currentHalfHour {
			continue
		}

		// 上传文件
		filePath := filepath.Join(u.logDir, filename)
		if err := u.uploadFile(filePath, logType); err != nil {
			logger.Error("上传日志文件失败", "file", filename, "error", err)
			failedCount++
		} else {
			// 上传成功，删除本地文件
			if err := os.Remove(filePath); err != nil {
				logger.Error("删除日志文件失败", "file", filename, "error", err)
			} else {
				logger.Info("日志文件已上传并删除", "file", filename)
				uploadedCount++
			}
		}
	}

	if uploadedCount > 0 || failedCount > 0 {
		logger.Info("日志上传完成", "uploaded", uploadedCount, "failed", failedCount)
	}
}

// parseLogFilename 解析日志文件名
// 返回：日志类型、半小时标识、是否解析成功
func parseLogFilename(filename string) (logType string, halfHour string, ok bool) {
	// 匹配格式: request-2025-01-15_10.log 或 system-2025-01-15_10.log
	re := regexp.MustCompile(`^(request|system)-(\d{4}-\d{2}-\d{2}_\d{2})\.log$`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) != 3 {
		return "", "", false
	}
	return matches[1], matches[2], true
}

// uploadFile 上传单个日志文件
func (u *Uploader) uploadFile(filePath, logType string) error {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 创建 multipart 表单
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件字段
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("创建表单文件失败: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("读取文件内容失败: %w", err)
	}

	// 添加类型字段
	if err := writer.WriteField("type", logType); err != nil {
		return fmt.Errorf("写入类型字段失败: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("关闭表单写入器失败: %w", err)
	}

	// 构建请求
	url := strings.TrimRight(u.serviceURL, "/") + "/api/logs/upload"
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+u.apiKey)

	// 发送请求
	resp, err := u.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 读取并丢弃响应体，以便连接可以被重用
		io.Copy(io.Discard, resp.Body)
		return fmt.Errorf("服务器返回错误: %s", resp.Status)
	}

	return nil
}

// ParseUploadInterval 解析上传间隔配置字符串（如 "30m", "1h"）
func ParseUploadInterval(s string) time.Duration {
	if s == "" {
		return uploadInterval
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return uploadInterval
	}
	return d
}

// FormatHalfHour 格式化半小时标识为可读时间
func FormatHalfHour(halfHourStr string) string {
	// 解析格式: 2025-01-15_10
	parts := strings.Split(halfHourStr, "_")
	if len(parts) != 2 {
		return halfHourStr
	}

	dateStr := parts[0]
	halfHourNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return halfHourStr
	}

	hour := halfHourNum / 2
	min := (halfHourNum % 2) * 30

	return fmt.Sprintf("%s %02d:%02d", dateStr, hour, min)
}
