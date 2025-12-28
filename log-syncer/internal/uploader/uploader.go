// Package uploader HTTP 上传器
package uploader

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Uploader 上传器
type Uploader struct {
	baseURL        string
	systemAuthToken string
	timeout        time.Duration
}

// NewUploader 创建上传器
func NewUploader(baseURL, systemAuthToken string, timeout time.Duration) *Uploader {
	return &Uploader{
		baseURL:        baseURL,
		systemAuthToken: systemAuthToken,
		timeout:        timeout,
	}
}

// UploadRequestLogs 上传请求日志
func (u *Uploader) UploadRequestLogs(entries []interface{}) error {
	if len(entries) == 0 {
		return nil
	}

	url := fmt.Sprintf("%s/api/request-logs/batch", u.baseURL)
	return u.upload(url, entries)
}

// UploadSystemLogs 上传系统日志
func (u *Uploader) UploadSystemLogs(entries []interface{}) error {
	if len(entries) == 0 {
		return nil
	}

	url := fmt.Sprintf("%s/api/system-logs/batch", u.baseURL)
	return u.upload(url, entries)
}

// upload 执行上传
func (u *Uploader) upload(url string, entries []interface{}) error {
	body, err := json.Marshal(entries)
	if err != nil {
		return fmt.Errorf("序列化数据失败: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", u.systemAuthToken))

	client := &http.Client{
		Timeout: u.timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("上传失败 (status=%d): %s", resp.StatusCode, string(respBody))
	}

	// 解析响应
	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Count int `json:"count"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析响应失败: %w", err)
	}

	if result.Code != 0 {
		return fmt.Errorf("上传失败: %s", result.Message)
	}

	log.Printf("上传成功: %d 条记录", result.Data.Count)
	return nil
}

// UploadResult 上传结果
type UploadResult struct {
	Success bool
	Error   error
}

// UploadWithRetry 上传（带重试，当前不重试）
func (u *Uploader) UploadWithRetry(entries []interface{}, isRequest bool) UploadResult {
	var err error
	if isRequest {
		err = u.UploadRequestLogs(entries)
	} else {
		err = u.UploadSystemLogs(entries)
	}

	return UploadResult{
		Success: err == nil,
		Error:   err,
	}
}
