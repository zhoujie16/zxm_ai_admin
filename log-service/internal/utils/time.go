// Package utils 工具函数包
// 提供通用的工具函数
package utils

import (
	"time"
)

// ParseTime 解析时间字符串
// 支持 RFC3339、RFC3339Nano 和带毫秒的 ISO 8601 格式
// 如果解析失败或为空，返回当前时间
func ParseTime(timeStr string) time.Time {
	if timeStr == "" {
		return time.Unix(0, 0)
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

	// 尝试带毫秒的 ISO 8601 格式：2006-01-02T15:04:05.000-07:00
	t, err = time.Parse("2006-01-02T15:04:05.000-07:00", timeStr)
	if err == nil {
		return t
	}

	// 解析失败时返回 1970-01-01 00:00:00 +08:00
	return time.Date(1970, 1, 1, 0, 0, 0, 0, time.FixedZone("UTC+8", 8*3600))
}

