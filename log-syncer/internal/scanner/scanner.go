// Package scanner 日志文件扫描器
package scanner

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
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

// LogFile 日志文件信息
type LogFile struct {
	Path     string    // 文件完整路径
	Name     string    // 文件名
	Type     LogType   // 日志类型
	FileTime time.Time // 文件名中的时间
}

// Scanner 扫描器
type Scanner struct {
	logDir string
}

// NewScanner 创建扫描器
func NewScanner(logDir string) *Scanner {
	return &Scanner{
		logDir: logDir,
	}
}

// Scan 扫描日志目录
// 只返回文件时间早于 cutoffTime 的文件
func (s *Scanner) Scan(cutoffTime time.Time) ([]*LogFile, error) {
	var files []*LogFile

	entries, err := os.ReadDir(s.logDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		logFile, ok := s.parseLogFile(entry.Name())
		if !ok {
			continue
		}

		// 只处理文件时间早于截止时间的文件
		if logFile.FileTime.Before(cutoffTime) || logFile.FileTime.Equal(cutoffTime) {
			files = append(files, logFile)
		}
	}

	return files, nil
}

// parseLogFile 解析日志文件名
// 格式: request-YYYYMMDDHHMM.log 或 system-YYYYMMDDHHMM.log
func (s *Scanner) parseLogFile(filename string) (*LogFile, bool) {
	// 匹配 request-YYYYMMDDHHMM.log 或 system-YYYYMMDDHHMM.log
	re := regexp.MustCompile(`^(request|system)-(\d{12})\.log$`)
	matches := re.FindStringSubmatch(filename)
	if matches == nil {
		return nil, false
	}

	// 解析时间（使用本地时区）
	fileTime, err := time.ParseInLocation("200601021504", matches[2], time.Local)
	if err != nil {
		return nil, false
	}

	logType := LogTypeRequest
	if matches[1] == "system" {
		logType = LogTypeSystem
	}

	return &LogFile{
		Name:     filename,
		Path:     filepath.Join(s.logDir, filename),
		Type:     logType,
		FileTime: fileTime,
	}, true
}

// ParseFileTime 从文件名解析时间 (YYYYMMDDHHMM)
func ParseFileTime(filename string) (time.Time, bool) {
	re := regexp.MustCompile(`^(\d{12})`)
	matches := re.FindStringSubmatch(filename)
	if matches == nil {
		return time.Time{}, false
	}

	t, err := time.ParseInLocation("200601021504", matches[1], time.Local)
	if err != nil {
		return time.Time{}, false
	}

	return t, true
}

// FormatFileTime 格式化时间为文件名时间格式 (YYYYMMDDHHMM)
func FormatFileTime(t time.Time) string {
	return t.Format("200601021504")
}

// GetArchiveRetentionTime 获取归档保留截止时间
func GetArchiveRetentionTime(retentionDays int) time.Time {
	return time.Now().AddDate(0, 0, -retentionDays)
}

// ShouldCleanArchive 判断归档文件是否应该清理
func ShouldCleanArchive(filename string, retentionDays int) bool {
	t, ok := ParseFileTime(filename)
	if !ok {
		return false
	}

	cutoffTime := GetArchiveRetentionTime(retentionDays)
	return t.Before(cutoffTime)
}

// StrToInt 字符串转整数
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
