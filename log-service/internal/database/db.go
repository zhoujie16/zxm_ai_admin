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

	// 修复 token_usage_logs 表的 request_id UNIQUE 约束
	if err := fixTokenUsageLogsUniqueConstraint(); err != nil {
		return fmt.Errorf("修复 token_usage_logs UNIQUE 约束失败: %w", err)
	}

	return nil
}

// fixTokenUsageLogsUniqueConstraint 修复 token_usage_logs 表的 request_id UNIQUE 约束
func fixTokenUsageLogsUniqueConstraint() error {
	// 检查是否已存在 UNIQUE 索引
	var count int64
	if err := DB.Raw(`
		SELECT COUNT(*) FROM sqlite_master 
		WHERE type='index' 
		AND name='idx_token_usage_logs_request_id' 
		AND sql LIKE '%UNIQUE%'
	`).Scan(&count).Error; err != nil {
		return err
	}

	// 如果不存在 UNIQUE 索引，则创建
	if count == 0 {
		// 删除旧的普通索引
		if err := DB.Exec(`DROP INDEX IF EXISTS idx_token_usage_logs_request_id`).Error; err != nil {
			return err
		}

		// 创建 UNIQUE 索引（等同于 UNIQUE 约束）
		if err := DB.Exec(`CREATE UNIQUE INDEX idx_token_usage_logs_request_id ON token_usage_logs(request_id)`).Error; err != nil {
			return err
		}

		logger.Info("已为 token_usage_logs 表添加 request_id UNIQUE 约束")
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
