package config

import (
	"log/slog"
	"os"
	"strings"
)

// Config 代理配置
type Config struct {
	TargetURL          string   // 目标转发地址
	OverrideAuthToken  string   // 替换的 Authorization 头
	RequiredAuthTokens []string // Authorization 白名单
	LogLevel           string   // 日志级别
	ListenAddr         string   // 监听地址
}

// Load 从环境变量加载配置
func Load() *Config {
	return &Config{
		TargetURL:          getEnv("TARGET_URL", "https://open.bigmodel.cn"),
		OverrideAuthToken:  getEnv("OVERRIDE_AUTH_TOKEN", "Bearer xxxx"),
		RequiredAuthTokens: splitTokens(getEnv("REQUIRED_AUTH_TOKENS", "Bearer 123456,Bearer abcdef")),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		ListenAddr:         getEnv("LISTEN_ADDR", ":6800"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func splitTokens(s string) []string {
	if s == "" {
		return nil
	}
	tokens := strings.Split(s, ",")
	result := make([]string, 0, len(tokens))
	for _, t := range tokens {
		trimmed := strings.TrimSpace(t)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
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
