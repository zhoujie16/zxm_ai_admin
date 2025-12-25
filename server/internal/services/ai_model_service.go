// Package services 业务逻辑服务层
// 实现 AI 模型管理相关的业务逻辑，包括增删改查操作
package services

import (
	"errors"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type AIModelService struct{}

// NewAIModelService 创建 AI 模型业务逻辑实例
func NewAIModelService() *AIModelService {
	return &AIModelService{}
}

// CreateAIModelRequest 创建 AI 模型请求
type CreateAIModelRequest struct {
	ModelKey  string `json:"model_key" binding:"required"` // 模型Key
	ModelName string `json:"model_name" binding:"required"` // 模型名称
	ApiURL    string `json:"api_url" binding:"required"`   // API地址
	Remark    string `json:"remark"`                       // 备注
	Status    int    `json:"status"`                       // 状态：1=启用，0=禁用
}

// UpdateAIModelRequest 更新 AI 模型请求
type UpdateAIModelRequest struct {
	ModelKey  *string `json:"model_key"`  // 模型Key
	ModelName *string `json:"model_name"` // 模型名称
	ApiURL    *string `json:"api_url"`    // API地址
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
	Total int64            `json:"total"` // 总数量
	List  []models.AIModel `json:"list"`  // 列表数据
}

// CreateAIModel 创建 AI 模型
func (s *AIModelService) CreateAIModel(req *CreateAIModelRequest) (*models.AIModel, error) {
	// 设置默认状态
	status := req.Status
	if status != 0 && status != 1 {
		status = 1 // 默认启用
	}

	// 创建 AI 模型
	aiModel := &models.AIModel{
		ModelKey:  req.ModelKey,
		ModelName: req.ModelName,
		ApiURL:    req.ApiURL,
		Remark:    req.Remark,
		Status:    status,
	}

	if err := database.DB.Create(aiModel).Error; err != nil {
		return nil, errors.New("创建 AI 模型失败")
	}

	return aiModel, nil
}

// GetAIModel 根据 ID 获取 AI 模型
func (s *AIModelService) GetAIModel(id uint) (*models.AIModel, error) {
	var aiModel models.AIModel
	if err := database.DB.First(&aiModel, id).Error; err != nil {
		return nil, errors.New("AI 模型不存在")
	}
	return &aiModel, nil
}

// ListAIModels 获取 AI 模型列表
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

	// 查询总数
	if err := database.DB.Model(&models.AIModel{}).Count(&total).Error; err != nil {
		return nil, errors.New("查询 AI 模型列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := database.DB.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询 AI 模型列表失败")
	}

	return &ListAIModelsResponse{
		Total: total,
		List:  list,
	}, nil
}

// UpdateAIModel 更新 AI 模型
func (s *AIModelService) UpdateAIModel(id uint, req *UpdateAIModelRequest) (*models.AIModel, error) {
	// 查询 AI 模型是否存在
	var aiModel models.AIModel
	if err := database.DB.First(&aiModel, id).Error; err != nil {
		return nil, errors.New("AI 模型不存在")
	}

	// 更新模型Key
	if req.ModelKey != nil {
		aiModel.ModelKey = *req.ModelKey
	}

	// 更新模型名称
	if req.ModelName != nil {
		aiModel.ModelName = *req.ModelName
	}

	// 更新API地址
	if req.ApiURL != nil {
		aiModel.ApiURL = *req.ApiURL
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
		return nil, errors.New("更新 AI 模型失败")
	}

	return &aiModel, nil
}

// DeleteAIModel 删除 AI 模型
func (s *AIModelService) DeleteAIModel(id uint) error {
	// 检查是否存在
	var aiModel models.AIModel
	if err := database.DB.First(&aiModel, id).Error; err != nil {
		return errors.New("AI 模型不存在")
	}

	// 物理删除
	if err := database.DB.Delete(&aiModel).Error; err != nil {
		return errors.New("删除 AI 模型失败")
	}

	return nil
}
