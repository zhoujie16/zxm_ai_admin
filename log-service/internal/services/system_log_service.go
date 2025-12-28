// Package services 业务逻辑服务层
// 实现系统日志相关的业务逻辑
package services

import (
	"errors"
	"time"
	"zxm_ai_admin/log-service/internal/config"
	"zxm_ai_admin/log-service/internal/database"
	"zxm_ai_admin/log-service/internal/models"

	"gorm.io/gorm/clause"
)

type SystemLogService struct{}

// NewSystemLogService 创建系统日志业务逻辑实例
func NewSystemLogService() *SystemLogService {
	return &SystemLogService{}
}

// CreateSystemLogRequest 创建系统日志请求
type CreateSystemLogRequest struct {
	RequestID string    `json:"request_id"`
	Time      time.Time `json:"time"`
	Level     string    `json:"level"`
	Msg       string    `json:"msg"`
}

// BatchCreateSystemLogsRequest 批量创建系统日志请求
type BatchCreateSystemLogsRequest struct {
	Logs []CreateSystemLogRequest `json:"logs" binding:"required"`
}

// CreateSystemLog 创建系统日志记录
func (s *SystemLogService) CreateSystemLog(req *CreateSystemLogRequest) (*models.SystemLog, error) {
	log := &models.SystemLog{
		RequestID: req.RequestID,
		Time:      req.Time,
		Level:     req.Level,
		Msg:       req.Msg,
	}

	if err := database.DB.Create(log).Error; err != nil {
		return nil, errors.New("创建系统日志记录失败")
	}

	return log, nil
}

// BatchCreateSystemLogs 批量创建系统日志记录（忽略重复 request_id）
func (s *SystemLogService) BatchCreateSystemLogs(reqs []CreateSystemLogRequest) (int, error) {
	if len(reqs) == 0 {
		return 0, nil
	}

	logs := make([]models.SystemLog, 0, len(reqs))
	for _, req := range reqs {
		logs = append(logs, models.SystemLog{
			RequestID: req.RequestID,
			Time:      req.Time,
			Level:     req.Level,
			Msg:       req.Msg,
		})
	}

	// 使用 OnConflict 忽略重复的 request_id
	if err := database.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "request_id"}},
		DoNothing: true,
	}).Create(&logs).Error; err != nil {
		return 0, errors.New("批量创建系统日志记录失败")
	}

	return len(logs), nil
}

// ListSystemLogsRequest 系统日志列表查询请求
type ListSystemLogsRequest struct {
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
	Level     string `form:"level"`
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
}

// ListSystemLogsResponse 系统日志列表查询响应
type ListSystemLogsResponse struct {
	Total int64               `json:"total"`
	List  []models.SystemLog `json:"list"`
}

// ListSystemLogs 获取系统日志列表
func (s *SystemLogService) ListSystemLogs(req *ListSystemLogsRequest) (*ListSystemLogsResponse, error) {
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
	var list []models.SystemLog

	query := database.DB.Model(&models.SystemLog{})

	if req.Level != "" {
		query = query.Where("level = ?", req.Level)
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

	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("查询系统日志列表失败")
	}

	offset := (page - 1) * pageSize
	if err := query.
		Order("time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询系统日志列表失败")
	}

	return &ListSystemLogsResponse{
		Total: total,
		List:  list,
	}, nil
}

// GetSystemLog 根据 ID 获取系统日志记录
func (s *SystemLogService) GetSystemLog(id uint) (*models.SystemLog, error) {
	var log models.SystemLog
	if err := database.DB.First(&log, id).Error; err != nil {
		return nil, errors.New("记录不存在")
	}
	return &log, nil
}

// DeleteSystemLogsByTimeRangeRequest 按时间范围删除系统日志
type DeleteSystemLogsByTimeRangeRequest struct {
	StartTime       string `json:"start_time" binding:"required"`
	EndTime         string `json:"end_time" binding:"required"`
	SystemAuthToken string `json:"system_auth_token" binding:"required"`
}

// DeleteSystemLogsByTimeRangeResponse 按时间范围删除系统日志响应
type DeleteSystemLogsByTimeRangeResponse struct {
	DeletedCount int64 `json:"deleted_count"`
}

// DeleteSystemLogsByTimeRange 按时间范围删除系统日志
func (s *SystemLogService) DeleteSystemLogsByTimeRange(req *DeleteSystemLogsByTimeRangeRequest) (*DeleteSystemLogsByTimeRangeResponse, error) {
	cfg := config.GetConfig()

	// 验证 System Auth Token
	if req.SystemAuthToken != cfg.API.SystemAuthToken {
		return nil, errors.New("无效的系统认证令牌")
	}

	// 解析时间范围
	startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
	if err != nil {
		return nil, errors.New("开始时间格式错误，正确格式为: 2006-01-02 15:04:05")
	}

	endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
	if err != nil {
		return nil, errors.New("结束时间格式错误，正确格式为: 2006-01-02 15:04:05")
	}

	if startTime.After(endTime) {
		return nil, errors.New("开始时间不能晚于结束时间")
	}

	// 执行硬删除
	result := database.DB.Unscoped().Where("time >= ? AND time <= ?", startTime, endTime).Delete(&models.SystemLog{})
	if result.Error != nil {
		return nil, errors.New("删除系统日志记录失败")
	}

	return &DeleteSystemLogsByTimeRangeResponse{
		DeletedCount: result.RowsAffected,
	}, nil
}
