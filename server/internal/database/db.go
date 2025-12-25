// Package database 数据库连接管理模块
// 负责初始化SQLite数据库连接，配置连接池，预留数据库迁移功能
package database

import (
	"fmt"
	"log"
	"time"

	"zxm_ai_admin/server/internal/config"
	"zxm_ai_admin/server/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("配置未初始化")
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.Database.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
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

	log.Println("数据库连接成功")
	return nil
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	if err := DB.AutoMigrate(
		&models.ProxyService{},
	); err != nil {
		return fmt.Errorf("迁移代理服务表失败: %w", err)
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

