// Package models 数据模型定义
// 定义 Token 使用记录的数据模型结构
package models

import (
	"time"

	"gorm.io/gorm"
)

// TokenUsageLog Token 使用记录模型
type TokenUsageLog struct {
	ID                uint           `json:"id" gorm:"primaryKey"`
	Time              time.Time      `json:"time" gorm:"not null;index"`                        // 时间戳
	Level             string         `json:"level" gorm:"size:20"`                              // 日志级别
	Msg               string         `json:"msg" gorm:"size:200"`                               // 日志消息
	RequestID         string         `json:"request_id" gorm:"size:64;index"`                   // 请求唯一标识
	Method            string         `json:"method" gorm:"size:10;index"`                       // HTTP 方法
	Path              string         `json:"path" gorm:"size:500;index"`                        // 请求路径
	Query             string         `json:"query" gorm:"size:1000"`                            // URL 查询参数
	RemoteAddr        string         `json:"remote_addr" gorm:"size:100"`                       // 客户端地址
	UserAgent         string         `json:"user_agent" gorm:"size:500"`                        // User-Agent
	XForwardedFor     string         `json:"x_forwarded_for" gorm:"size:100"`                   // X-Forwarded-For
	RequestHeaders    map[string]string `json:"request_headers" gorm:"type:json"`              // 请求头（JSON）
	Authorization     string         `json:"authorization" gorm:"size:500;index"`               // Authorization 头
	RequestBody       string         `json:"request_body" gorm:"type:text"`                     // 请求体
	Status            int            `json:"status" gorm:"index"`                               // HTTP 响应状态码
	ResponseHeaders   map[string]string `json:"response_headers" gorm:"type:json"`             // 响应头（JSON）
	LatencyMs         int64          `json:"latency_ms"`                                       // 请求耗时（毫秒）
	RequestSizeBytes  int            `json:"request_size_bytes"`                               // 请求体大小
	ResponseSizeBytes int            `json:"response_size_bytes"`                              // 响应体大小
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (TokenUsageLog) TableName() string {
	return "token_usage_logs"
}
