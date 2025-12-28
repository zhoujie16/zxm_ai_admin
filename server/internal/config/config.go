// Package config 配置管理模块
// 负责加载和解析YAML配置文件，提供全局配置访问接口
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Server          ServerConfig   `mapstructure:"server"`
	Database        DatabaseConfig `mapstructure:"database"`
	Admin           AdminConfig    `mapstructure:"admin"`
	JWT             JWTConfig      `mapstructure:"jwt"`
	Log             LogConfig      `mapstructure:"log"`
	SystemAuthToken string         `mapstructure:"system_auth_token"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	Path         string `mapstructure:"path"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type AdminConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Dir   string `mapstructure:"dir"`
}

var AppConfig *Config

// Load 加载配置文件
func Load(configPath string) error {
	// 设置配置文件路径
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	// 设置环境变量支持
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 解析配置到结构体
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	// 确保数据库目录存在
	dbDir := filepath.Dir(AppConfig.Database.Path)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("创建数据库目录失败: %w", err)
	}

	return nil
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	return AppConfig
}

