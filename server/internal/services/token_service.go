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
	Total int64                  `json:"total"` // 总数量
	List  []TokenWithModelName    `json:"list"`  // 列表数据
}

// TokenWithModelName 包含模型名称的 Token 数据
type TokenWithModelName struct {
	ID           uint       `json:"id"`
	Token        string     `json:"token"`
	AIModelID    uint       `json:"ai_model_id"`
	ModelName    string     `json:"model_name"`    // 关联的模型名称
	OrderNo      string     `json:"order_no"`
	Status       int        `json:"status"`
	ExpireAt     *time.Time `json:"expire_at"`
	UsageLimit   int        `json:"usage_limit"`
	Remark       string     `json:"remark"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// TokenWithFullModel 包含完整模型信息的 Token 数据（字段带前缀）
type TokenWithFullModel struct {
	TokenID          uint       `json:"token_id"`
	Token            string     `json:"token"`
	TokenOrderNo     string     `json:"token_order_no"`
	TokenStatus      int        `json:"token_status"`
	TokenExpireAt    *time.Time `json:"token_expire_at"`
	TokenUsageLimit  int        `json:"token_usage_limit"`
	TokenRemark      string     `json:"token_remark"`
	AIModelID        uint       `json:"ai_model_id"`
	AIModelName      string     `json:"ai_model_name"`
	AIModelApiURL    string     `json:"ai_model_api_url"`
	AIModelApiKey    string     `json:"ai_model_api_key"`
	AIModelRemark    string     `json:"ai_model_remark"`
	AIModelStatus    int        `json:"ai_model_status"`
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
	var list []TokenWithModelName

	// 构建查询，关联 AI 模型表
	query := database.DB.Table("tokens t").
		Select(`t.id, t.token, t.ai_model_id, m.model_name,
			t.order_no, t.status, t.expire_at, t.usage_limit,
			t.remark, t.created_at, t.updated_at`).
		Joins("LEFT JOIN ai_models m ON t.ai_model_id = m.id").
		Where("t.deleted_at IS NULL")

	// 关键词搜索
	if req.Keyword != "" {
		query = query.Where("t.token LIKE ? OR t.remark LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 查询总数
	if err := database.DB.Model(&models.Token{}).Count(&total).Error; err != nil {
		return nil, errors.New("查询 Token 列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := query.
		Order("t.created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&list).Error; err != nil {
		return nil, errors.New("查询 Token 列表失败")
	}

	return &ListTokensResponse{
		Total: total,
		List:  list,
	}, nil
}

// ListAllTokensWithModel 获取所有 Token 及其完整模型信息（不分页）
func (s *TokenService) ListAllTokensWithModel() ([]TokenWithFullModel, error) {
	var list []TokenWithFullModel

	// 构建查询，关联 AI 模型表（INNER JOIN 过滤掉模型已不存在的 Token）
	if err := database.DB.Table("tokens t").
		Select(`t.id as token_id, t.token, t.order_no as token_order_no,
			t.status as token_status, t.expire_at as token_expire_at,
			t.usage_limit as token_usage_limit, t.remark as token_remark,
			t.ai_model_id,
			m.model_name as ai_model_name, m.api_url as ai_model_api_url,
			m.api_key as ai_model_api_key, m.remark as ai_model_remark,
			m.status as ai_model_status`).
		Joins("INNER JOIN ai_models m ON t.ai_model_id = m.id").
		Where("t.deleted_at IS NULL").
		Where("(t.expire_at IS NULL OR t.expire_at > ?)", time.Now()).
		Order("t.created_at DESC").
		Scan(&list).Error; err != nil {
		return nil, errors.New("查询 Token 列表失败")
	}

	return list, nil
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

// ListRecycledTokensRequest 回收站列表查询请求
type ListRecycledTokensRequest struct {
	Page     int    `form:"page"`      // 页码，从1开始
	PageSize int    `form:"page_size"` // 每页数量
	Keyword  string `form:"keyword"`   // 关键词搜索
}

// ListRecycledTokensResponse 回收站列表查询响应
type ListRecycledTokensResponse struct {
	Total int64                      `json:"total"` // 总数量
	List  []TokenWithModelName       `json:"list"`  // 列表数据
}

// ListRecycledTokens 获取已删除的 Token 列表
func (s *TokenService) ListRecycledTokens(req *ListRecycledTokensRequest) (*ListRecycledTokensResponse, error) {
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
	var list []TokenWithModelName

	// 构建查询，只查询已删除的记录
	query := database.DB.Table("tokens t").
		Select(`t.id, t.token, t.ai_model_id, m.model_name,
			t.order_no, t.status, t.expire_at, t.usage_limit,
			t.remark, t.created_at, t.updated_at`).
		Joins("LEFT JOIN ai_models m ON t.ai_model_id = m.id").
		Where("t.deleted_at IS NOT NULL")

	// 关键词搜索
	if req.Keyword != "" {
		query = query.Where("t.token LIKE ? OR t.remark LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	// 查询总数
	if err := database.DB.Table("tokens t").Where("t.deleted_at IS NOT NULL").Count(&total).Error; err != nil {
		return nil, errors.New("查询回收站列表失败")
	}

	// 查询列表
	offset := (page - 1) * pageSize
	if err := query.
		Order("t.deleted_at DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&list).Error; err != nil {
		return nil, errors.New("查询回收站列表失败")
	}

	return &ListRecycledTokensResponse{
		Total: total,
		List:  list,
	}, nil
}

// RestoreToken 恢复已删除的 Token
func (s *TokenService) RestoreToken(id uint) error {
	// 使用 Unscoped() 查询包括已删除的记录
	var token models.Token
	if err := database.DB.Unscoped().First(&token, id).Error; err != nil {
		return errors.New("Token 不存在")
	}

	// 如果没有被删除，不需要恢复
	if token.DeletedAt.Time.IsZero() {
		return errors.New("Token 未被删除，无需恢复")
	}

	// 恢复（将 deleted_at 设置为 NULL）
	if err := database.DB.Unscoped().Model(&token).Update("deleted_at", nil).Error; err != nil {
		return errors.New("恢复 Token 失败")
	}

	return nil
}

// DestroyToken 永久删除 Token
func (s *TokenService) DestroyToken(id uint) error {
	// 使用 Unscoped() 查询包括已删除的记录
	var token models.Token
	if err := database.DB.Unscoped().First(&token, id).Error; err != nil {
		return errors.New("Token 不存在")
	}

	// 永久删除
	if err := database.DB.Unscoped().Delete(&token).Error; err != nil {
		return errors.New("永久删除 Token 失败")
	}

	return nil
}
