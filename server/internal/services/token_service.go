// Package services 业务逻辑服务层
// 实现 Token 管理相关的业务逻辑，包括增删改查操作
package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/models"
)

type TokenService struct{}

// NewTokenService 创建 Token 业务逻辑实例
func NewTokenService() *TokenService {
	return &TokenService{}
}

// CreateTokenRequest 创建 Token 请求
type CreateTokenRequest struct {
	AIModelID  uint       `json:"ai_model_id" binding:"required"` // 关联的AI模型ID
	OrderNo    string     `json:"order_no"`                       // 关联订单号
	Status     int        `json:"status"`                         // 状态：1=启用，0=禁用
	ExpireAt   *time.Time `json:"expire_at"`                     // 过期时间
	UsageLimit int        `json:"usage_limit"`                   // 使用限额
	Remark     string     `json:"remark"`                        // 备注
}

// UpdateTokenRequest 更新 Token 请求
type UpdateTokenRequest struct {
	AIModelID  *uint       `json:"ai_model_id"`  // 关联的AI模型ID
	OrderNo    *string     `json:"order_no"`     // 关联订单号
	Status     *int        `json:"status"`       // 状态：1=启用，0=禁用
	ExpireAt   *time.Time  `json:"expire_at"`    // 过期时间
	UsageLimit *int        `json:"usage_limit"`  // 使用限额
	Remark     *string     `json:"remark"`       // 备注
}

// ListTokensRequest 列表查询请求
type ListTokensRequest struct {
	Page     int    `form:"page"`      // 页码，从1开始
	PageSize int    `form:"page_size"` // 每页数量
	Keyword  string `form:"keyword"`   // 关键词搜索（token或备注）
}

// ListTokensResponse 列表查询响应
type ListTokensResponse struct {
	Total int64        `json:"total"` // 总数量
	List  []models.Token `json:"list"`  // 列表数据
}

// CreateToken 创建 Token
func (s *TokenService) CreateToken(req *CreateTokenRequest) (*models.Token, error) {
	// 生成随机 Token
	tokenStr, err := GenerateRandomToken()
	if err != nil {
		return nil, errors.New("生成 Token 失败")
	}

	// 检查关联的 AI 模型是否存在
	var aiModel models.AIModel
	if err := database.DB.First(&aiModel, req.AIModelID).Error; err != nil {
		return nil, errors.New("关联的 AI 模型不存在")
	}

	// 设置默认状态
	status := req.Status
	if status != 0 && status != 1 {
		status = 1 // 默认启用
	}

	// 创建 Token
	token := &models.Token{
		Token:      tokenStr,
		AIModelID:  req.AIModelID,
		OrderNo:    req.OrderNo,
		Status:     status,
		ExpireAt:   req.ExpireAt,
		UsageLimit: req.UsageLimit,
		Remark:     req.Remark,
	}

	if err := database.DB.Create(token).Error; err != nil {
		return nil, errors.New("创建 Token 失败")
	}

	return token, nil
}

// GetToken 根据 ID 获取 Token
func (s *TokenService) GetToken(id uint) (*models.Token, error) {
	var token models.Token
	if err := database.DB.First(&token, id).Error; err != nil {
		return nil, errors.New("Token 不存在")
	}
	return &token, nil
}

// ListTokens 获取 Token 列表
func (s *TokenService) ListTokens(req *ListTokensRequest) (*ListTokensResponse, error) {
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
	var list []models.Token

	// 构建查询
	query := database.DB.Model(&models.Token{})

	// 关键词搜索
	if req.Keyword != "" {
		query = query.Where("token LIKE ? OR remark LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 查询总数
	if err := query.Count(&total).Error; err != nil {
		return nil, errors.New("查询 Token 列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := query.
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, errors.New("查询 Token 列表失败")
	}

	return &ListTokensResponse{
		Total: total,
		List:  list,
	}, nil
}

// UpdateToken 更新 Token
func (s *TokenService) UpdateToken(id uint, req *UpdateTokenRequest) (*models.Token, error) {
	// 查询 Token 是否存在
	var token models.Token
	if err := database.DB.First(&token, id).Error; err != nil {
		return nil, errors.New("Token 不存在")
	}

	// 更新关联模型
	if req.AIModelID != nil {
		// 检查关联的 AI 模型是否存在
		var aiModel models.AIModel
		if err := database.DB.First(&aiModel, *req.AIModelID).Error; err != nil {
			return nil, errors.New("关联的 AI 模型不存在")
		}
		token.AIModelID = *req.AIModelID
	}

	// 更新订单号
	if req.OrderNo != nil {
		token.OrderNo = *req.OrderNo
	}

	// 更新状态
	if req.Status != nil {
		if *req.Status != 0 && *req.Status != 1 {
			return nil, errors.New("状态值无效，只能为0或1")
		}
		token.Status = *req.Status
	}

	// 更新过期时间
	if req.ExpireAt != nil {
		token.ExpireAt = req.ExpireAt
	}

	// 更新使用限额
	if req.UsageLimit != nil {
		if *req.UsageLimit < 0 {
			return nil, errors.New("使用限额不能为负数")
		}
		token.UsageLimit = *req.UsageLimit
	}

	// 更新备注
	if req.Remark != nil {
		token.Remark = *req.Remark
	}

	// 保存更新
	if err := database.DB.Save(&token).Error; err != nil {
		return nil, errors.New("更新 Token 失败")
	}

	return &token, nil
}

// DeleteToken 删除 Token（软删除）
func (s *TokenService) DeleteToken(id uint) error {
	// 检查是否存在
	var token models.Token
	if err := database.DB.First(&token, id).Error; err != nil {
		return errors.New("Token 不存在")
	}

	// 软删除
	if err := database.DB.Delete(&token).Error; err != nil {
		return errors.New("删除 Token 失败")
	}

	return nil
}

// GenerateRandomToken 生成随机 Token
func GenerateRandomToken() (string, error) {
	// 生成 32 字节的随机数，转成 64 个十六进制字符
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return "sk-" + hex.EncodeToString(bytes), nil
}
