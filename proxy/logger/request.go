package logger

import (
	"context"
	"log/slog"
)

// RequestLogger 请求日志记录器
type RequestLogger struct {
	baseLogger
}

var request = &RequestLogger{}

// Init 初始化请求日志
func (r *RequestLogger) Init(logDir string, level slog.Level) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.init(logDir, level, "request-")
}

// Log 记录请求日志（带 context）
func (r *RequestLogger) Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	logger := r.getLogger()
	if logger == nil {
		return
	}

	r.rotate()
	logger.Log(ctx, level, msg, args...)
}

// Info 记录 info 级别日志
func (r *RequestLogger) Info(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	logger := r.getLogger()
	if logger == nil {
		return
	}

	r.rotate()
	logger.Info(msg, args...)
}

// Error 记录 error 级别日志
func (r *RequestLogger) Error(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	logger := r.getLogger()
	if logger == nil {
		return
	}

	r.rotate()
	logger.Error(msg, args...)
}

// Warn 记录 warn 级别日志
func (r *RequestLogger) Warn(msg string, args ...any) {
	r.mu.Lock()
	defer r.mu.Unlock()

	logger := r.getLogger()
	if logger == nil {
		return
	}

	r.rotate()
	logger.Warn(msg, args...)
}

// Request 导出的请求日志实例
var Request = request
