// Package models 数据模型定义
// 定义代理服务相关的数据模型结构
package models

import (
	"time"

	"gorm.io/gorm"
)

// ProxyService 代理服务模型
type ProxyService struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ServiceID string         `json:"service_id" gorm:"uniqueIndex;not null;size:100" binding:"required"` // 服务标识，唯一
	ServerIP  string         `json:"server_ip" gorm:"not null;size:50" binding:"required"`                // 服务器IP
	Status    int            `json:"status" gorm:"default:1"`                                            // 状态：1=启用，0=未启用
	Remark    string         `json:"remark" gorm:"size:500"`                                             // 备注
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 指定表名
func (ProxyService) TableName() string {
	return "proxy_services"
}


