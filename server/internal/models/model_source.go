// Package models 数据模型定义
// 定义模型来源相关的数据模型结构
package models

import (
	"time"

	"gorm.io/gorm"
)

// ModelSource 模型来源
type ModelSource struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ModelName string         `json:"model_name" gorm:"not null;size:100" binding:"required"` // 模型名称
	ApiURL    string         `json:"api_url" gorm:"not null;size:500" binding:"required"`    // API地址
	ApiKey    string         `json:"api_key" gorm:"uniqueIndex;not null;size:255" binding:"required"` // API Key，唯一
	Remark    string         `json:"remark" gorm:"size:500"` // 备注
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (ModelSource) TableName() string {
	return "model_sources"
}
