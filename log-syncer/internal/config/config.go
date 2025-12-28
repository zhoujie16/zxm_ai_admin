// Package config 配置管理
package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	cfg  *Config
	once sync.Once
)

// Config 配置结构
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Proxy    ProxyConfig    `yaml:"proxy"`
	Archive  ArchiveConfig  `yaml:"archive"`
	Uploader UploaderConfig `yaml:"uploader"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	LogServiceURL   string        `yaml:"log_service_url"`
	SystemAuthToken string        `yaml:"system_auth_token"`
	Timeout         time.Duration `yaml:"timeout"`
}

// ProxyConfig Proxy 日志配置
type ProxyConfig struct {
	LogDir string `yaml:"log_dir"`
}

// ArchiveConfig 归档配置
type ArchiveConfig struct {
	Dir            string `yaml:"dir"`
	RetentionDays  int    `yaml:"retention_days"`
}

// UploaderConfig 上传配置
type UploaderConfig struct {
	BatchSize int `yaml:"batch_size"`
}

// Load 加载配置文件
func Load(path string) error {
	var err error
	once.Do(func() {
		cfg = &Config{}
		err = loadFile(path, cfg)
	})
	return err
}

// loadFile 从文件加载配置
func loadFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 设置默认值
	if cfg.Uploader.BatchSize == 0 {
		cfg.Uploader.BatchSize = 100
	}

	return nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return cfg
}
