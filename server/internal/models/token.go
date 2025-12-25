// Package models 数据模型定义
// 定义 Token 相关的数据模型结构
package models

import (
	"time"

	"gorm.io/gorm"
)

// Token Token 模型
type Token struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Token        string         `json:"token" gorm:"uniqueIndex;not null;size:255" binding:"required"` // Token值，唯一
	AIModelID    uint           `json:"ai_model_id" gorm:"not null;index"`                             // 关联的AI模型ID
	OrderNo      string         `json:"order_no" gorm:"size:100"`                                       // 关联订单号
	Status       int            `json:"status" gorm:"default:1"`                                        // 状态：1=启用，0=禁用
	ExpireAt     *time.Time     `json:"expire_at"`                                                      // 过期时间
	UsageLimit   int            `json:"usage_limit" gorm:"default:0"`                                   // 使用限额（调用次数）
	Remark       string         `json:"remark" gorm:"size:500"`                                         // 备注
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (Token) TableName() string {
	return "tokens"
}
