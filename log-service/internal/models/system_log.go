// Package models 数据模型定义
// 定义系统日志的数据模型结构
package models

import (
	"time"

	"gorm.io/gorm"
)

// SystemLog 系统日志模型
type SystemLog struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	RequestID string         `json:"request_id" gorm:"size:64;unique"`
	Time      time.Time      `json:"time" gorm:"not null;index"`
	Level     string         `json:"level" gorm:"size:20;index"`
	Msg       string         `json:"msg" gorm:"size:500"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (SystemLog) TableName() string {
	return "system_logs"
}
