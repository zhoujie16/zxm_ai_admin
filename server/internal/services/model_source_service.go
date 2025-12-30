// Package services 业务逻辑服务层
// 实现模型来源管理相关的业务逻辑，包括增删改查操作
package services

import (
	"errors"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type ModelSourceService struct{}

// NewModelSourceService 创建模型来源业务逻辑实例
func NewModelSourceService() *ModelSourceService {
	return &ModelSourceService{}
}

// CreateModelSourceRequest 创建模型来源请求
type CreateModelSourceRequest struct {
	ModelName string `json:"model_name" binding:"required"` // 模型名称
	ApiURL    string `json:"api_url" binding:"required"`    // API地址
	ApiKey    string `json:"api_key" binding:"required"`    // API Key
	Remark    string `json:"remark"`                        // 备注
}

// UpdateModelSourceRequest 更新模型来源请求
type UpdateModelSourceRequest struct {
	ModelName *string `json:"model_name"` // 模型名称
	Remark    *string `json:"remark"`     // 备注
}

// ListModelSourcesRequest 列表查询请求
type ListModelSourcesRequest struct {
	Page     int `form:"page"`      // 页码，从1开始
	PageSize int `form:"page_size"` // 每页数量
}

// ListModelSourcesResponse 列表查询响应
type ListModelSourcesResponse struct {
	Total int64               `json:"total"` // 总数量
	List  []models.ModelSource `json:"list"`  // 列表数据
}

// CreateModelSource 创建模型来源
func (s *ModelSourceService) CreateModelSource(req *CreateModelSourceRequest) (*models.ModelSource, error) {
	// 检查 API Key 是否已存在
	var count int64
	if err := database.DB.Model(&models.ModelSource{}).Where("api_key = ?", req.ApiKey).Count(&count).Error; err != nil {
		return nil, errors.New("检查 API Key 失败")
	}
	if count > 0 {
		return nil, errors.New("API Key 已存在")
	}

	// 创建模型来源
	modelSource := &models.ModelSource{
		ModelName: req.ModelName,
		ApiURL:    req.ApiURL,
		ApiKey:    req.ApiKey,
		Remark:    req.Remark,
	}

	if err := database.DB.Create(modelSource).Error; err != nil {
		return nil, errors.New("创建模型来源失败")
	}

	return modelSource, nil
}

// GetModelSource 根据 ID 获取模型来源
func (s *ModelSourceService) GetModelSource(id uint) (*models.ModelSource, error) {
	var modelSource models.ModelSource
	if err := database.DB.First(&modelSource, id).Error; err != nil {
		return nil, errors.New("模型来源不存在")
	}
	return &modelSource, nil
}

// ListModelSources 获取模型来源列表
func (s *ModelSourceService) ListModelSources(req *ListModelSourcesRequest) (*ListModelSourcesResponse, error) {
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
	var list []models.ModelSource

	// 查询总数
	if err := database.DB.Model(&models.ModelSource{}).Count(&total).Error; err != nil {
		return nil, errors.New("查询模型来源列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := database.DB.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询模型来源列表失败")
	}

	return &ListModelSourcesResponse{
		Total: total,
		List:  list,
	}, nil
}

// UpdateModelSource 更新模型来源
func (s *ModelSourceService) UpdateModelSource(id uint, req *UpdateModelSourceRequest) (*models.ModelSource, error) {
	// 查询模型来源是否存在
	var modelSource models.ModelSource
	if err := database.DB.First(&modelSource, id).Error; err != nil {
		return nil, errors.New("模型来源不存在")
	}

	// 更新模型名称
	if req.ModelName != nil {
		modelSource.ModelName = *req.ModelName
	}

	// 更新备注
	if req.Remark != nil {
		modelSource.Remark = *req.Remark
	}

	// 保存更新
	if err := database.DB.Save(&modelSource).Error; err != nil {
		return nil, errors.New("更新模型来源失败")
	}

	return &modelSource, nil
}

// DeleteModelSource 删除模型来源
func (s *ModelSourceService) DeleteModelSource(id uint) error {
	// 检查是否存在
	var modelSource models.ModelSource
	if err := database.DB.First(&modelSource, id).Error; err != nil {
		return errors.New("模型来源不存在")
	}

	// 软删除
	if err := database.DB.Delete(&modelSource).Error; err != nil {
		return errors.New("删除模型来源失败")
	}

	return nil
}
