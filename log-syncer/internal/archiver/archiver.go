// Package archiver 归档与清理模块
package archiver

import (
	"log"
	"os"
	"path/filepath"
	"zxm_ai_admin/log-syncer/internal/scanner"
)

// Archiver 归档器
type Archiver struct {
	archiveDir     string
	retentionDays  int
}

// NewArchiver 创建归档器
func NewArchiver(archiveDir string, retentionDays int) *Archiver {
	return &Archiver{
		archiveDir:    archiveDir,
		retentionDays: retentionDays,
	}
}

// Archive 归档文件
func (a *Archiver) Archive(srcPath string) error {
	// 确保归档目录存在
	if err := os.MkdirAll(a.archiveDir, 0755); err != nil {
		return err
	}

	// 获取文件名
	filename := filepath.Base(srcPath)
	dstPath := filepath.Join(a.archiveDir, filename)

	// 移动文件
	if err := os.Rename(srcPath, dstPath); err != nil {
		return err
	}

	log.Printf("归档文件: %s → %s", srcPath, dstPath)
	return nil
}

// CleanExpired 清理过期的归档文件
func (a *Archiver) CleanExpired() error {
	if err := os.MkdirAll(a.archiveDir, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(a.archiveDir)
	if err != nil {
		return err
	}

	cleanedCount := 0
	cutoffTime := scanner.GetArchiveRetentionTime(a.retentionDays)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		if scanner.ShouldCleanArchive(filename, a.retentionDays) {
			path := filepath.Join(a.archiveDir, filename)
			if err := os.Remove(path); err == nil {
				log.Printf("删除过期归档: %s (时间: %s)", filename, cutoffTime.Format("2006-01-02 15:04:05"))
				cleanedCount++
			}
		}
	}

	if cleanedCount > 0 {
		log.Printf("清理过期归档完成: 删除 %d 个文件", cleanedCount)
	}

	return nil
}
