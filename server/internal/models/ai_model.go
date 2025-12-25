// Package models 数据模型定义
// 定义 AI 模型相关的数据模型结构
package models

import "time"

// AIModel AI 模型
type AIModel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ModelKey  string    `json:"model_key" gorm:"not null;size:100" binding:"required"`      // 模型Key
	ModelName string    `json:"model_name" gorm:"not null;size:100" binding:"required"`     // 模型名称
	ApiURL    string    `json:"api_url" gorm:"not null;size:500" binding:"required"`        // 模型调用API地址
	Remark    string    `json:"remark" gorm:"size:500"`                                      // 备注
	Status    int       `json:"status" gorm:"default:1"`                                     // 状态：1=启用，0=禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (AIModel) TableName() string {
	return "ai_models"
}
