// Package services 业务逻辑服务层
// 实现 Token 使用记录相关的业务逻辑，包括新增和查询操作
package services

import (
	"errors"
	"time"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type TokenUsageLogService struct{}

// NewTokenUsageLogService 创建 Token 使用记录业务逻辑实例
func NewTokenUsageLogService() *TokenUsageLogService {
	return &TokenUsageLogService{}
}

// CreateTokenUsageLogRequest 创建 Token 使用记录请求
type CreateTokenUsageLogRequest struct {
	Time              time.Time         `json:"time" binding:"required"`
	Level             string            `json:"level"`
	Msg               string            `json:"msg"`
	RequestID         string            `json:"request_id" binding:"required"`
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

// ListTokenUsageLogsRequest Token 使用记录列表查询请求
type ListTokenUsageLogsRequest struct {
	Page         int    `form:"page"`           // 页码，从1开始
	PageSize     int    `form:"page_size"`      // 每页数量
	RequestID    string `form:"request_id"`     // 按 request_id 精确查询
	StartTime    string `form:"start_time"`     // 开始时间（格式：2006-01-02 15:04:05）
	EndTime      string `form:"end_time"`       // 结束时间
	Status       int    `form:"status"`         // 按状态码过滤（-1 表示全部）
	Method       string `form:"method"`         // 按 HTTP 方法过滤
	Authorization string `form:"authorization"` // 按 Authorization 头模糊匹配
}

// ListTokenUsageLogsResponse Token 使用记录列表查询响应
type ListTokenUsageLogsResponse struct {
	Total int64                        `json:"total"` // 总数量
	List  []models.TokenUsageLog       `json:"list"`  // 列表数据
}

// CreateTokenUsageLog 创建 Token 使用记录
func (s *TokenUsageLogService) CreateTokenUsageLog(req *CreateTokenUsageLogRequest) (*models.TokenUsageLog, error) {
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
		return nil, errors.New("创建 Token 使用记录失败")
	}

	return log, nil
}

// ListTokenUsageLogs 获取 Token 使用记录列表
func (s *TokenUsageLogService) ListTokenUsageLogs(req *ListTokenUsageLogsRequest) (*ListTokenUsageLogsResponse, error) {
	// 设置默认分页参数
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

	// 构建查询
	query := database.DB.Model(&models.TokenUsageLog{})

	// 按 request_id 精确查询
	if req.RequestID != "" {
		query = query.Where("request_id = ?", req.RequestID)
	}

	// 按时间范围过滤
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

	// 按状态码过滤
	if req.Status >= 0 {
		query = query.Where("status = ?", req.Status)
	}

	// 按 HTTP 方法过滤
	if req.Method != "" {
		query = query.Where("method = ?", req.Method)
	}

	// 按 Authorization 头模糊匹配
	if req.Authorization != "" {
		query = query.Where("authorization LIKE ?", "%"+req.Authorization+"%")
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("查询 Token 使用记录列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := query.
		Order("time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询 Token 使用记录列表失败")
	}

	return &ListTokenUsageLogsResponse{
		Total: total,
		List:  list,
	}, nil
}

// GetTokenUsageLog 根据 ID 获取 Token 使用记录
func (s *TokenUsageLogService) GetTokenUsageLog(id uint) (*models.TokenUsageLog, error) {
	var log models.TokenUsageLog
	if err := database.DB.First(&log, id).Error; err != nil {
		return nil, errors.New("记录不存在")
	}
	return &log, nil
}
