package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Request 请求日志记录器
type RequestLogger struct {
	logDir  string
	level   slog.Level
	mu      sync.Mutex
	logger  *slog.Logger
	currentTimestamp string
	file    *os.File
}

var request = &RequestLogger{}

// Init 初始化请求日志
func (r *RequestLogger) Init(logDir string, level slog.Level) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logDir = logDir
	r.level = level

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	return r.rotate()
}

// rotate 切换日志文件（调用前必须已加锁）
func (r *RequestLogger) rotate() error {
	currentTimestamp := getCurrentHalfHour()

	// 如果半小时未变化，无需切换
	if r.currentTimestamp == currentTimestamp && r.logger != nil {
		return nil
	}

	// 关闭旧文件
	if r.file != nil {
		r.file.Close()
	}

	// 创建新文件
	filename := filepath.Join(r.logDir, "request-"+currentTimestamp+".log")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	r.file = file
	r.currentTimestamp = currentTimestamp

	// 创建新的 logger
	r.logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: r.level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().UTC().Format(time.RFC3339))
			}
			return a
		},
	}))

	return nil
}

// Log 记录请求日志（带 context）
func (r *RequestLogger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.logger == nil {
		return
	}

	r.rotate()
	r.logger.Log(ctx, level, msg, args...)
}

// Info 记录 info 级别日志
func (r *RequestLogger) Info(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.logger == nil {
		return
	}

	r.rotate()
	r.logger.Info(msg, args...)
}

// Error 记录 error 级别日志
func (r *RequestLogger) Error(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.logger == nil {
		return
	}

	r.rotate()
	r.logger.Error(msg, args...)
}

// Warn 记录 warn 级别日志
func (r *RequestLogger) Warn(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.logger == nil {
		return
	}

	r.rotate()
	r.logger.Warn(msg, args...)
}

// Request 导出的请求日志实例
var Request = request
