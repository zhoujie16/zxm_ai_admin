// Package logger 日志记录模块
// 提供便捷的日志记录函数和请求 ID 管理
package logger

import (
	"context"

	"github.com/google/uuid"
)

type contextKey struct{}

// GenerateRequestID 生成请求 ID（UUID 格式）
func GenerateRequestID() string {
	return uuid.New().String()
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

// Info 记录 info 日志（便捷函数，委托给 System.Info）
func Info(msg string, args ...any) {
	System.Info(msg, args...)
}

// Error 记录 error 日志（便捷函数，委托给 System.Error）
func Error(msg string, args ...any) {
	System.Error(msg, args...)
}

// Warn 记录 warn 日志（便捷函数，委托给 System.Warn）
func Warn(msg string, args ...any) {
	System.Warn(msg, args...)
}

// Debug 记录 debug 日志（便捷函数，委托给 System.Debug）
func Debug(msg string, args ...any) {
	System.Debug(msg, args...)
}

