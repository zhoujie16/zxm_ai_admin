// Package services 业务逻辑服务层
// 实现管理员认证相关的业务逻辑，包括登录验证和获取管理员信息
package services

import (
	"errors"

	"zxm_ai_admin/server/internal/config"
	"zxm_ai_admin/server/internal/utils"
)

type AuthService struct{}

// NewAuthService 创建认证服务实例
func NewAuthService() *AuthService {
	return &AuthService{}
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token    string      `json:"token"`
	Username string      `json:"username"`
	UserInfo interface{} `json:"user_info"`
}

// AdminInfo 管理员信息
type AdminInfo struct {
	Username string `json:"username"`
}

// Login 管理员登录
func (s *AuthService) Login(req *LoginRequest) (*LoginResponse, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil, errors.New("配置未初始化")
	}

	// 验证用户名和密码
	if req.Username != cfg.Admin.Username || req.Password != cfg.Admin.Password {
		return nil, errors.New("用户名或密码错误")
	}

	// 生成token（使用固定的用户ID 1）
	token, err := utils.GenerateToken(1, cfg.Admin.Username)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &LoginResponse{
		Token:    token,
		Username: cfg.Admin.Username,
		UserInfo: AdminInfo{
			Username: cfg.Admin.Username,
		},
	}, nil
}

// GetAdminInfo 获取管理员信息
func (s *AuthService) GetAdminInfo() (*AdminInfo, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return nil, errors.New("配置未初始化")
	}

	return &AdminInfo{
		Username: cfg.Admin.Username,
	}, nil
}

