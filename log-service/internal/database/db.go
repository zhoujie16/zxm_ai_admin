// Package database 数据库连接管理模块
// 负责初始化 SQLite 数据库连接，配置连接池，自动迁移表结构
package database

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"zxm_ai_admin/log-service/internal/config"
	"zxm_ai_admin/log-service/internal/logger"
	"zxm_ai_admin/log-service/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("配置未初始化")
	}

	// 确保数据库目录存在
	dbDir := filepath.Dir(cfg.Database.Path)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("创建数据库目录失败: %w", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}

	logger.Info("数据库连接成功", "path", cfg.Database.Path)
	return nil
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	if err := DB.AutoMigrate(
		&models.TokenUsageLog{},
		&models.SystemLog{},
	); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}
	return nil
}

// Close 关闭数据库连接
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
