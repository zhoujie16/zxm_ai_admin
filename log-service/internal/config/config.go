// Package config 配置管理模块
// 负责加载和管理应用配置
package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
	API      APIConfig      `yaml:"api"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Path          string `yaml:"path"`
	MaxOpenConns  int    `yaml:"max_open_conns"`
	MaxIdleConns  int    `yaml:"max_idle_conns"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level string `yaml:"level"` // 日志级别：debug, info, warn, error
	Dir   string `yaml:"dir"`   // 日志文件目录
}

// APIConfig API 配置
type APIConfig struct {
	SystemAuthToken string `yaml:"system_auth_token"` // proxy 写入日志用的系统认证令牌
	JWTSecret      string `yaml:"jwt_secret"`        // 用于验证 JWT
}

var (
	cfg  *Config
	once sync.Once
)

// Load 加载配置文件
func Load(configPath string) error {
	var err error
	once.Do(func() {
		data, err := os.ReadFile(configPath)
		if err != nil {
			err = fmt.Errorf("读取配置文件失败: %w", err)
			return
		}

		cfg = &Config{}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			err = fmt.Errorf("解析配置文件失败: %w", err)
			return
		}

		// 设置默认值
		if cfg.Server.Port == 0 {
			cfg.Server.Port = 6809
		}
		if cfg.Server.Mode == "" {
			cfg.Server.Mode = "release"
		}
		if cfg.Database.Path == "" {
			cfg.Database.Path = "./data/logs.db"
		}
		if cfg.Database.MaxOpenConns == 0 {
			cfg.Database.MaxOpenConns = 10
		}
		if cfg.Database.MaxIdleConns == 0 {
			cfg.Database.MaxIdleConns = 5
		}
		if cfg.Log.Level == "" {
			cfg.Log.Level = "info"
		}
		if cfg.Log.Dir == "" {
			cfg.Log.Dir = "./logs"
		}
	})

	return err
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return cfg
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
