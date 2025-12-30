// Package services 业务逻辑服务层
// 实现模型代理管理相关的业务逻辑，包括增删改查操作
package services

import (
	"errors"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type AIModelService struct{}

// NewAIModelService 创建模型代理业务逻辑实例
func NewAIModelService() *AIModelService {
	return &AIModelService{}
}

// CreateAIModelRequest 创建模型代理请求
type CreateAIModelRequest struct {
	ModelName string `json:"model_name" binding:"required"` // 模型名称
	ApiKey    string `json:"api_key" binding:"required"`    // 模型来源的API Key（用于查询模型来源）
	Remark    string `json:"remark"`                         // 备注
	Status    int    `json:"status"`                         // 状态：1=启用，0=禁用
}

// UpdateAIModelRequest 更新模型代理请求
type UpdateAIModelRequest struct {
	ModelName *string `json:"model_name"` // 模型名称
	ApiKey    *string `json:"api_key"`    // 模型来源的API Key（如需更换API信息）
	Remark    *string `json:"remark"`     // 备注
	Status    *int    `json:"status"`     // 状态：1=启用，0=禁用
}

// ListAIModelsRequest 列表查询请求
type ListAIModelsRequest struct {
	Page     int `form:"page"`      // 页码，从1开始
	PageSize int `form:"page_size"` // 每页数量
}

// ListAIModelsResponse 列表查询响应
type ListAIModelsResponse struct {
	Total int64           `json:"total"` // 总数量
	List  []models.AIModel `json:"list"`  // 列表数据
}

// CreateAIModel 创建模型代理
func (s *AIModelService) CreateAIModel(req *CreateAIModelRequest) (*models.AIModel, error) {
	// 设置默认状态
	status := req.Status
	if status != 0 && status != 1 {
		status = 1 // 默认启用
	}

	// 根据 api_key 查询模型来源，获取 API 地址
	var modelSource models.ModelSource
	if err := database.DB.Where("api_key = ?", req.ApiKey).First(&modelSource).Error; err != nil {
		return nil, errors.New("模型来源不存在")
	}

	// 创建模型代理，复制 API 信息
	aiModel := &models.AIModel{
		ModelName: req.ModelName,
		ApiURL:    modelSource.ApiURL,
		ApiKey:    modelSource.ApiKey,
		Remark:    req.Remark,
		Status:    status,
	}

	if err := database.DB.Create(aiModel).Error; err != nil {
		return nil, errors.New("创建模型代理失败")
	}

	return aiModel, nil
}

// GetAIModel 根据 ID 获取模型代理
func (s *AIModelService) GetAIModel(id uint) (*models.AIModel, error) {
	var aiModel models.AIModel
	if err := database.DB.Unscoped().First(&aiModel, id).Error; err != nil {
		return nil, errors.New("模型代理不存在")
	}
	return &aiModel, nil
}

// ListAIModels 获取模型代理列表
func (s *AIModelService) ListAIModels(req *ListAIModelsRequest) (*ListAIModelsResponse, error) {
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
	var list []models.AIModel

	// 查询总数（排除软删除）
	if err := database.DB.Model(&models.AIModel{}).Count(&total).Error; err != nil {
		return nil, errors.New("查询模型代理列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := database.DB.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询模型代理列表失败")
	}

	return &ListAIModelsResponse{
		Total: total,
		List:  list,
	}, nil
}

// UpdateAIModel 更新模型代理
func (s *AIModelService) UpdateAIModel(id uint, req *UpdateAIModelRequest) (*models.AIModel, error) {
	// 查询模型代理是否存在
	var aiModel models.AIModel
	if err := database.DB.Unscoped().First(&aiModel, id).Error; err != nil {
		return nil, errors.New("模型代理不存在")
	}

	// 更新模型名称
	if req.ModelName != nil {
		aiModel.ModelName = *req.ModelName
	}

	// 如果提供了 api_key，重新获取 API 信息
	if req.ApiKey != nil {
		var modelSource models.ModelSource
		if err := database.DB.Where("api_key = ?", *req.ApiKey).First(&modelSource).Error; err != nil {
			return nil, errors.New("模型来源不存在")
		}
		aiModel.ApiURL = modelSource.ApiURL
		aiModel.ApiKey = modelSource.ApiKey
	}

	// 更新状态
	if req.Status != nil {
		if *req.Status != 0 && *req.Status != 1 {
			return nil, errors.New("状态值无效，只能为0或1")
		}
		aiModel.Status = *req.Status
	}

	// 更新备注
	if req.Remark != nil {
		aiModel.Remark = *req.Remark
	}

	// 保存更新
	if err := database.DB.Save(&aiModel).Error; err != nil {
		return nil, errors.New("更新模型代理失败")
	}

	return &aiModel, nil
}

// DeleteAIModel 删除模型代理
func (s *AIModelService) DeleteAIModel(id uint) error {
	// 检查是否存在
	var aiModel models.AIModel
	if err := database.DB.First(&aiModel, id).Error; err != nil {
		return errors.New("模型代理不存在")
	}

	// 软删除
	if err := database.DB.Delete(&aiModel).Error; err != nil {
		return errors.New("删除模型代理失败")
	}

	return nil
}
