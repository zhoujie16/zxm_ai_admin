// Package models 数据模型定义
// 定义 Token 使用记录相关的数据模型结构
package models

import "time"

// TokenUsageLog Token 使用记录
type TokenUsageLog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Token     string    `json:"token" gorm:"not null;size:255;index"` // Token值
	RemoteIP  string    `json:"remote_ip" gorm:"size:100"`              // 请求来源IP
	UserAgent string    `json:"user_agent" gorm:"size:500"`            // User Agent
	CallTime  time.Time `json:"call_time" gorm:"not null;index"`       // 调用时间
	CreatedAt time.Time `json:"created_at"`
}

// TableName 指定表名
func (TokenUsageLog) TableName() string {
	return "token_usage_logs"
}
