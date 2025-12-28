// Package models 数据模型定义
// 定义 Token 使用记录的数据模型结构
package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// JSONMap 用于存储 map[string]string 到数据库的 JSON 类型
type JSONMap map[string]string

// Scan 实现 sql.Scanner 接口
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONMap value")
	}
	return json.Unmarshal(bytes, j)
}

// Value 实现 driver.Valuer 接口
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// TokenUsageLog Token 使用记录模型
type TokenUsageLog struct {
	ID                uint            `json:"id" gorm:"primaryKey"`
	Time              time.Time       `json:"time" gorm:"not null;index"`
	Level             string          `json:"level" gorm:"size:20"`
	Msg               string          `json:"msg" gorm:"size:200"`
	RequestID         string          `json:"request_id" gorm:"size:64;index"`
	Method            string          `json:"method" gorm:"size:10;index"`
	Path              string          `json:"path" gorm:"size:500;index"`
	Query             string          `json:"query" gorm:"size:1000"`
	RemoteAddr        string          `json:"remote_addr" gorm:"size:100"`
	UserAgent         string          `json:"user_agent" gorm:"size:500"`
	XForwardedFor     string          `json:"x_forwarded_for" gorm:"size:100"`
	RequestHeaders    JSONMap         `json:"request_headers" gorm:"type:text"`
	Authorization     string          `json:"authorization" gorm:"size:500;index"`
	RequestBody       string          `json:"request_body" gorm:"type:text"`
	Status            int             `json:"status" gorm:"index"`
	ResponseHeaders   JSONMap         `json:"response_headers" gorm:"type:text"`
	LatencyMs         int64           `json:"latency_ms"`
	RequestSizeBytes  int             `json:"request_size_bytes"`
	ResponseSizeBytes int             `json:"response_size_bytes"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `json:"-" gorm:"index"`
}

// TableName 指定表名
func (TokenUsageLog) TableName() string {
	return "token_usage_logs"
}
