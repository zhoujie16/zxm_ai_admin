package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// baseLogger 日志记录器基类，包含公共字段和方法
type baseLogger struct {
	logDir          string
	level           slog.Level
	mu              sync.Mutex
	logger          *slog.Logger
	currentTimestamp string
	file            *os.File
	filePrefix      string // 文件名前缀，如 "request-" 或 "system-"
}

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
	// 处理跨天情况
	date := now
	if hour >= 24 {
		date = now.AddDate(0, 0, 1)
		hour = 0
	}
	return date.Format("20060102") + fmt.Sprintf("%02d%02d", hour, minEnd)
}

// init 初始化日志记录器（调用前必须已加锁）
func (b *baseLogger) init(logDir string, level slog.Level, filePrefix string) error {
	b.logDir = logDir
	b.level = level
	b.filePrefix = filePrefix

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	return b.rotate()
}

// rotate 切换日志文件（调用前必须已加锁）
func (b *baseLogger) rotate() error {
	currentTimestamp := getCurrentHalfHour()

	// 如果半小时未变化，无需切换
	if b.currentTimestamp == currentTimestamp && b.logger != nil {
		return nil
	}

	// 关闭旧文件
	if b.file != nil {
		if err := b.file.Close(); err != nil {
			// 记录关闭错误，但不影响日志切换流程
			// 注意：这里不能使用 logger，因为可能正在初始化
			_ = err // 忽略错误，避免影响日志切换
		}
	}

	// 创建新文件
	filename := filepath.Join(b.logDir, b.filePrefix+currentTimestamp+".log")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	b.file = file
	b.currentTimestamp = currentTimestamp

	// 创建新的 logger
	b.logger = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: b.level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().Format("2006-01-02T15:04:05.000-07:00"))
			}
			return a
		},
	}))

	return nil
}

// getLogger 获取 logger 实例（调用前必须已加锁）
func (b *baseLogger) getLogger() *slog.Logger {
	return b.logger
}

