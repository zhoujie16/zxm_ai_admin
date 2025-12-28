package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

// Config 代理配置
type Config struct {
	LogLevel           string `mapstructure:"log_level"`
	ListenAddr         string `mapstructure:"listen_addr"`
	ServerBaseURL      string `mapstructure:"server_base_url"`
	ServerUsername     string `mapstructure:"server_username"`
	ServerPassword     string `mapstructure:"server_password"`
	SyncInterval       int    `mapstructure:"sync_interval"`
}

var appConfig *Config

// Load 加载配置文件
func Load(configPath string) (*Config, error) {
	v := viper.New()

	// 设置配置文件路径
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	cfg := &Config{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	appConfig = cfg
	return cfg, nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return appConfig
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
