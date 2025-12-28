package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

// System 系统日志记录器
type SystemLogger struct {
	logDir  string
	level   slog.Level
	mu      sync.Mutex
	logger  *slog.Logger
	currentTimestamp string
	file    *os.File
}

var system = &SystemLogger{}

// getCurrentHalfHour 获取当前半小时标识（用于文件名）
// 返回紧凑时间戳格式，如：202512281630（表示 2025-12-28 16:00-16:30 的半小时）
func getCurrentHalfHour() string {
	now := time.Now()
	minute := now.Minute()
	// 计算当前半小时的结束时间
	hour := now.Hour()
	minEnd := ((minute / 30) + 1) * 30
	// 如果是 60 分钟，进位到下一小时
	if minEnd >= 60 {
		hour += 1
		minEnd = 0
	}
	return now.Format("20060102") + fmt.Sprintf("%02d%02d", hour, minEnd)
}

// Init 初始化系统日志
func (s *SystemLogger) Init(logDir string, level slog.Level) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.logDir = logDir
	s.level = level

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	return s.rotate()
}

// rotate 切换日志文件（调用前必须已加锁）
func (s *SystemLogger) rotate() error {
	currentTimestamp := getCurrentHalfHour()

	// 如果半小时未变化，无需切换
	if s.currentTimestamp == currentTimestamp && s.logger != nil {
		return nil
	}

	// 关闭旧文件
	if s.file != nil {
		s.file.Close()
	}

	// 创建新文件
	filename := filepath.Join(s.logDir, "system-"+currentTimestamp+".log")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	s.file = file
	s.currentTimestamp = currentTimestamp

	// 创建新的 logger
	s.logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: s.level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().UTC().Format(time.RFC3339))
			}
			return a
		},
	}))

	return nil
}

// Info 记录 info 级别日志
func (s *SystemLogger) Info(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	s.logger.Info(msg, newArgs...)
}

// Error 记录 error 级别日志
func (s *SystemLogger) Error(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	s.logger.Error(msg, newArgs...)
}

// Warn 记录 warn 级别日志
func (s *SystemLogger) Warn(msg string, args ...any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.logger == nil {
		return
	}

	s.rotate()
	requestID := uuid.New().String()
	newArgs := append([]any{"request_id", requestID}, args...)
	s.logger.Warn(msg, newArgs...)
}

// System 导出的系统日志实例
var System = system
