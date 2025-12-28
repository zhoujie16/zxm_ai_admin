// Package utils 工具函数包
// 提供通用的工具函数
package utils

import (
	"time"
)

// ParseTime 解析时间字符串
// 支持 RFC3339 和 RFC3339Nano 格式
// 如果解析失败或为空，返回当前时间
func ParseTime(timeStr string) time.Time {
	if timeStr == "" {
		return time.Now()
	}

	// 尝试 RFC3339 格式
	t, err := time.Parse(time.RFC3339, timeStr)
	if err == nil {
		return t
	}

	// 尝试 RFC3339Nano 格式
	t, err = time.Parse(time.RFC3339Nano, timeStr)
	if err == nil {
		return t
	}

	// 解析失败时返回当前时间
	return time.Now()
}

