package logger

import (
	"context"
	"log/slog"
	"math/rand"
	"os"
)

type contextKey struct{}

// Init 初始化日志
func Init(level slog.Level) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)
	return logger
}

// GenerateRequestID 生成请求 ID
func GenerateRequestID() string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// ContextWithRequestID 将 requestID 存入 context
func ContextWithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, contextKey{}, requestID)
}

// RequestIDFromContext 从 context 获取 requestID
func RequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(contextKey{}).(string); ok {
		return requestID
	}
	return ""
}

// LogRequest 记录请求日志
func LogRequest(ctx context.Context, level slog.Level, msg string, args ...any) {
	slog.Default().Log(ctx, level, msg, args...)
}

// Info 记录 info 日志
func Info(msg string, args ...any) {
	slog.Default().Info(msg, args...)
}

// Error 记录 error 日志
func Error(msg string, args ...any) {
	slog.Default().Error(msg, args...)
}

// Warn 记录 warn 日志
func Warn(msg string, args ...any) {
	slog.Default().Warn(msg, args...)
}
