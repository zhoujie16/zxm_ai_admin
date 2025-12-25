// Package services 业务逻辑服务层
// 实现 Token 使用记录相关的业务逻辑
package services

import (
	"errors"
	"time"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type TokenUsageService struct{}

// NewTokenUsageService 创建 Token 使用记录业务逻辑实例
func NewTokenUsageService() *TokenUsageService {
	return &TokenUsageService{}
}

// RecordUsageRequest 记录使用请求
type RecordUsageRequest struct {
	Token     string `json:"token" binding:"required"`     // Token值
	RemoteIP  string `json:"remote_ip" binding:"required"` // 请求来源IP
	UserAgent string `json:"user_agent"`                   // User Agent
}

// ListUsageLogsRequest 列表查询请求
type ListUsageLogsRequest struct {
	Page     int    `form:"page"`      // 页码，从1开始
	PageSize int    `form:"page_size"` // 每页数量
	Token    string `form:"token"`     // Token值（筛选）
}

// ListUsageLogsResponse 列表查询响应
type ListUsageLogsResponse struct {
	Total int64                 `json:"total"` // 总数量
	List  []models.TokenUsageLog `json:"list"`  // 列表数据
}

// RecordUsage 记录 Token 使用
func (s *TokenUsageService) RecordUsage(req *RecordUsageRequest) error {
	// 检查 Token 是否存在
	var token models.Token
	if err := database.DB.Where("token = ?", req.Token).First(&token).Error; err != nil {
		return errors.New("Token 不存在")
	}

	// 创建使用记录
	usageLog := &models.TokenUsageLog{
		Token:     req.Token,
		RemoteIP:  req.RemoteIP,
		UserAgent: req.UserAgent,
		CallTime:  time.Now(),
	}

	if err := database.DB.Create(usageLog).Error; err != nil {
		return errors.New("记录使用失败")
	}

	return nil
}

// GetUsageCount 获取 Token 使用次数
func (s *TokenUsageService) GetUsageCount(token string) (int64, error) {
	var count int64
	if err := database.DB.Model(&models.TokenUsageLog{}).
		Where("token = ?", token).
		Count(&count).Error; err != nil {
		return 0, errors.New("查询使用次数失败")
	}
	return count, nil
}

// ListUsageLogs 获取使用记录列表
func (s *TokenUsageService) ListUsageLogs(req *ListUsageLogsRequest) (*ListUsageLogsResponse, error) {
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

	// Token 筛选
	if req.Token != "" {
		query = query.Where("token = ?", req.Token)
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("查询使用记录列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := query.
		Order("call_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询使用记录列表失败")
	}

	return &ListUsageLogsResponse{
		Total: total,
		List:  list,
	}, nil
}
