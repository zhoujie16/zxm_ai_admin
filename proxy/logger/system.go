package logger

import (
	"log/slog"

	"github.com/google/uuid"
)

// SystemLogger 系统日志记录器
type SystemLogger struct {
	baseLogger
}

var system = &SystemLogger{}

// Init 初始化系统日志
func (s *SystemLogger) Init(logDir string, level slog.Level) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.init(logDir, level, "system-")
}

// Info 记录 info 级别日志
func (s *SystemLogger) Info(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	logger := s.getLogger()
	if logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	logger.Info(msg, newArgs...)
}

// Error 记录 error 级别日志
func (s *SystemLogger) Error(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	logger := s.getLogger()
	if logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	logger.Error(msg, newArgs...)
}

// Warn 记录 warn 级别日志
func (s *SystemLogger) Warn(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	logger := s.getLogger()
	if logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	logger.Warn(msg, newArgs...)
}

// System 导出的系统日志实例
var System = system
