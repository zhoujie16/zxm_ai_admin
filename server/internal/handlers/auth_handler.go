// Package handlers HTTP请求处理器
// 处理认证相关的HTTP请求，包括管理员登录和获取当前管理员信息
package handlers

import (
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler 创建认证处理器实例
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录获取token
// @Tags 认证
// @Accept json
// @Produce json
// @Param body body services.LoginRequest true "登录信息"
// @Success 200 {object} utils.Response
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.authService.Login(&req)
	if err != nil {
		utils.Error(c, 401, err.Error())
		return
	}

	utils.Success(c, response)
}

// GetMe 获取当前管理员信息
// @Summary 获取当前管理员信息
// @Description 获取当前登录管理员的信息
// @Tags 认证
// @Security BearerAuth
// @Produce json
// @Success 200 {object} utils.Response
// @Router /api/auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	adminInfo, err := h.authService.GetAdminInfo()
	if err != nil {
		utils.Error(c, 500, err.Error())
		return
	}

	utils.Success(c, adminInfo)
}

