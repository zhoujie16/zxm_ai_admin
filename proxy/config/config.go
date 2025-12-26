package config

import (
	"log/slog"
	"os"
	"strings"
)

// Config 代理配置
type Config struct {
	LogLevel   string // 日志级别
	ListenAddr string // 监听地址

	// Token 动态路由配置
	ServerAPIURL   string // 后端 API 地址
	ServerAPIToken string // 后端 API 认证 Token
	SyncInterval   int    // 缓存同步间隔（分钟）
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		ListenAddr:     getEnv("LISTEN_ADDR", ":6800"),
		ServerAPIURL:   getEnv("SERVER_API_URL", "http://localhost:6808"),
		ServerAPIToken: getEnv("SERVER_API_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoienhtLWFpLWFkbWluIiwiZXhwIjoxNzY2NzU5NTQzLCJuYmYiOjE3NjY2NzMxNDMsImlhdCI6MTc2NjY3MzE0M30.mOyW65CHdokcH5OJks5DU3pKHT24-RtLpuK8Nb7nFzs"),
		SyncInterval:   getEnvInt("SYNC_INTERVAL", 10),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := parseInt(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func parseInt(s string) (int, error) {
	var result int
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0, &parseError{s}
		}
		result = result*10 + int(ch-'0')
	}
	return result, nil
}

// parseError 解析错误
type parseError struct {
	input string
}

func (e *parseError) Error() string {
	return "invalid integer: " + e.input
}

// ParseLogLevel 解析日志级别
func ParseLogLevel(level string) slog.Level {
	switch strings.ToLower(level) {
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
