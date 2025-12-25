// Package handlers HTTP请求处理器
// 处理 Token 管理相关的 HTTP 请求
package handlers

import (
	"net/http"
	"strconv"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenService *services.TokenService
}

// NewTokenHandler 创建 Token 处理器实例
func NewTokenHandler() *TokenHandler {
	return &TokenHandler{
		tokenService: services.NewTokenService(),
	}
}

// CreateToken 创建 Token
func (h *TokenHandler) CreateToken(c *gin.Context) {
	var req services.CreateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	token, err := h.tokenService.CreateToken(&req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, token)
}

// GetToken 获取 Token 详情
func (h *TokenHandler) GetToken(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	token, err := h.tokenService.GetToken(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, token)
}

// ListTokens 获取 Token 列表
func (h *TokenHandler) ListTokens(c *gin.Context) {
	var req services.ListTokensRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.tokenService.ListTokens(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// UpdateToken 更新 Token
func (h *TokenHandler) UpdateToken(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	var req services.UpdateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	token, err := h.tokenService.UpdateToken(uint(id), &req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, token)
}

// DeleteToken 删除 Token
func (h *TokenHandler) DeleteToken(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	if err := h.tokenService.DeleteToken(uint(id)); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, nil)
}
