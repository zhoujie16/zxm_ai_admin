// Package parser 日志解析器
package parser

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// LogType 日志类型
type LogType int

const (
	// LogTypeRequest 请求日志
	LogTypeRequest LogType = iota
	// LogTypeSystem 系统日志
	LogTypeSystem
)

// LogEntry 日志条目接口
type LogEntry interface{}

// RequestLogEntry 请求日志条目
type RequestLogEntry map[string]interface{}

// SystemLogEntry 系统日志条目
type SystemLogEntry struct {
	RequestID string `json:"request_id"`
	Time      string `json:"time"`
	Level     string `json:"level"`
	Msg       string `json:"msg"`
}

// Parser 解析器
type Parser struct{}

// NewParser 创建解析器
func NewParser() *Parser {
	return &Parser{}
}

// ParseFile 解析日志文件
func (p *Parser) ParseFile(filePath string, logType LogType) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	var entries []LogEntry
	logScanner := bufio.NewScanner(file)

	// 增加缓冲区大小以处理长日志行
	const maxScanTokenSize = 1024 * 1024 // 1MB
	buf := make([]byte, 0, maxScanTokenSize)
	logScanner.Buffer(buf, maxScanTokenSize)

	for logScanner.Scan() {
		line := logScanner.Text()
		if line == "" {
			continue
		}

		if logType == LogTypeRequest {
			entry, err := p.parseRequestLog(line)
			if err != nil {
				continue // 跳过解析失败的行
			}
			entries = append(entries, entry)
		} else {
			entry, err := p.parseSystemLog(line)
			if err != nil {
				continue // 跳过解析失败的行
			}
			entries = append(entries, entry)
		}
	}

	if err := logScanner.Err(); err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	return entries, nil
}

// parseRequestLog 解析请求日志
func (p *Parser) parseRequestLog(line string) (RequestLogEntry, error) {
	var entry map[string]interface{}
	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		return nil, err
	}

	// 验证必填字段
	if _, ok := entry["request_id"]; !ok {
		return nil, errors.New("缺少 request_id 字段")
	}

	return entry, nil
}

// parseSystemLog 解析系统日志
func (p *Parser) parseSystemLog(line string) (*SystemLogEntry, error) {
	var entry map[string]interface{}
	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		return nil, err
	}

	// 验证必填字段
	requestID, ok := entry["request_id"].(string)
	if !ok || requestID == "" {
		return nil, errors.New("缺少 request_id 字段")
	}

	// 提取需要的字段
	timeStr, _ := entry["time"].(string)
	level, _ := entry["level"].(string)
	msg, _ := entry["msg"].(string)

	// 设置默认值
	if timeStr == "" {
		timeStr = time.Now().Format("2006-01-02T15:04:05.000-07:00")
	}
	if level == "" {
		level = "INFO"
	}

	return &SystemLogEntry{
		RequestID: requestID,
		Time:      timeStr,
		Level:     level,
		Msg:       msg,
	}, nil
}

// Batch 分批处理日志条目
func (p *Parser) Batch(entries []LogEntry, batchSize int) [][]LogEntry {
	var batches [][]LogEntry

	for i := 0; i < len(entries); i += batchSize {
		end := i + batchSize
		if end > len(entries) {
			end = len(entries)
		}
		batches = append(batches, entries[i:end])
	}

	return batches
}
