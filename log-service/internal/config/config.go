// Package config 配置管理模块
// 负责加载和管理应用配置
package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
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
	})

	return err
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return cfg
}
