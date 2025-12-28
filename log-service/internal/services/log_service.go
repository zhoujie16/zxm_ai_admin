// Package services 业务逻辑服务层
// 实现 Token 使用记录相关的业务逻辑，包括新增和查询操作
package services

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"zxm_ai_admin/log-service/internal/database"
	"zxm_ai_admin/log-service/internal/models"
)

type LogService struct{}

// NewLogService 创建日志业务逻辑实例
func NewLogService() *LogService {
	return &LogService{}
}

// CreateLogRequest 创建日志请求
type CreateLogRequest struct {
	Time              time.Time         `json:"time"`
	Level             string            `json:"level"`
	Msg               string            `json:"msg"`
	RequestID         string            `json:"request_id"`
	Method            string            `json:"method"`
	Path              string            `json:"path"`
	Query             string            `json:"query"`
	RemoteAddr        string            `json:"remote_addr"`
	UserAgent         string            `json:"user_agent"`
	XForwardedFor     string            `json:"x_forwarded_for"`
	RequestHeaders    map[string]string `json:"request_headers"`
	Authorization     string            `json:"authorization"`
	RequestBody       string            `json:"request_body"`
	Status            int               `json:"status"`
	ResponseHeaders   map[string]string `json:"response_headers"`
	LatencyMs         int64             `json:"latency_ms"`
	RequestSizeBytes  int               `json:"request_size_bytes"`
	ResponseSizeBytes int               `json:"response_size_bytes"`
}

// ListLogsRequest 日志列表查询请求
type ListLogsRequest struct {
	Page          int    `form:"page"`
	PageSize      int    `form:"page_size"`
	RequestID     string `form:"request_id"`
	StartTime     string `form:"start_time"`
	EndTime       string `form:"end_time"`
	Status        string `form:"status"`
	Method        string `form:"method"`
	Authorization string `form:"authorization"`
}

// ListLogsResponse 日志列表查询响应
type ListLogsResponse struct {
	Total int64                   `json:"total"`
	List  []models.TokenUsageLog `json:"list"`
}

// CreateLog 创建日志记录
func (s *LogService) CreateLog(req *CreateLogRequest) (*models.TokenUsageLog, error) {
	log := &models.TokenUsageLog{
		Time:              req.Time,
		Level:             req.Level,
		Msg:               req.Msg,
		RequestID:         req.RequestID,
		Method:            req.Method,
		Path:              req.Path,
		Query:             req.Query,
		RemoteAddr:        req.RemoteAddr,
		UserAgent:         req.UserAgent,
		XForwardedFor:     req.XForwardedFor,
		RequestHeaders:    req.RequestHeaders,
		Authorization:     req.Authorization,
		RequestBody:       req.RequestBody,
		Status:            req.Status,
		ResponseHeaders:   req.ResponseHeaders,
		LatencyMs:         req.LatencyMs,
		RequestSizeBytes:  req.RequestSizeBytes,
		ResponseSizeBytes: req.ResponseSizeBytes,
	}

	if err := database.DB.Create(log).Error; err != nil {
		return nil, errors.New("创建日志记录失败")
	}

	return log, nil
}

// ListLogs 获取日志列表
func (s *LogService) ListLogs(req *ListLogsRequest) (*ListLogsResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	var total int64
	var list []models.TokenUsageLog

	query := database.DB.Model(&models.TokenUsageLog{})

	if req.RequestID != "" {
		query = query.Where("request_id = ?", req.RequestID)
	}

	if req.StartTime != "" {
		if startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime); err == nil {
			query = query.Where("time >= ?", startTime)
		}
	}
	if req.EndTime != "" {
		if endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime); err == nil {
			query = query.Where("time <= ?", endTime)
		}
	}

	// 解析状态码字符串（逗号分隔，支持单个或多个）
	if req.Status != "" {
		parts := strings.Split(req.Status, ",")
		var statuses []int
		for _, part := range parts {
			if num, err := strconv.Atoi(strings.TrimSpace(part)); err == nil {
				statuses = append(statuses, num)
			}
		}
		if len(statuses) > 0 {
			query = query.Where("status IN ?", statuses)
		}
	}

	if req.Method != "" {
		query = query.Where("method = ?", req.Method)
	}

	if req.Authorization != "" {
		query = query.Where("authorization LIKE ?", "%"+req.Authorization+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("查询日志列表失败")
	}

	offset := (page - 1) * pageSize
	if err := query.
		Order("time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询日志列表失败")
	}

	return &ListLogsResponse{
		Total: total,
		List:  list,
	}, nil
}

// GetLog 根据 ID 获取日志记录
func (s *LogService) GetLog(id uint) (*models.TokenUsageLog, error) {
	var log models.TokenUsageLog
	if err := database.DB.First(&log, id).Error; err != nil {
		return nil, errors.New("记录不存在")
	}
	return &log, nil
}

// BatchCreateLogs 批量创建日志记录
func (s *LogService) BatchCreateLogs(reqs []CreateLogRequest) int {
	if len(reqs) == 0 {
		return 0
	}

	logs := make([]models.TokenUsageLog, 0, len(reqs))
	for _, req := range reqs {
		logs = append(logs, models.TokenUsageLog{
			Time:              req.Time,
			Level:             req.Level,
			Msg:               req.Msg,
			RequestID:         req.RequestID,
			Method:            req.Method,
			Path:              req.Path,
			Query:             req.Query,
			RemoteAddr:        req.RemoteAddr,
			UserAgent:         req.UserAgent,
			XForwardedFor:     req.XForwardedFor,
			RequestHeaders:    req.RequestHeaders,
			Authorization:     req.Authorization,
			RequestBody:       req.RequestBody,
			Status:            req.Status,
			ResponseHeaders:   req.ResponseHeaders,
			LatencyMs:         req.LatencyMs,
			RequestSizeBytes:  req.RequestSizeBytes,
			ResponseSizeBytes: req.ResponseSizeBytes,
		})
	}

	if err := database.DB.Create(&logs).Error; err != nil {
		return 0
	}

	return len(logs)
}
