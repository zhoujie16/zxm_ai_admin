// Package models 数据模型定义
// 定义模型代理相关的数据模型结构
package models

import "time"

// AIModel 模型代理
type AIModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ModelName string    `json:"model_name" gorm:"not null;size:100" binding:"required"` // 模型名称
	ApiURL    string    `json:"api_url" gorm:"not null;size:500"`                        // API地址
	ApiKey    string    `json:"api_key" gorm:"size:255"`                                  // API Key
	Remark    string    `json:"remark" gorm:"size:500"`                                   // 备注
	Status    int       `json:"status" gorm:"default:1"`                                  // 状态：1=启用，0=禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"` // 软删除
}

// TableName 指定表名
func (AIModel) TableName() string {
	return "ai_models"
}
