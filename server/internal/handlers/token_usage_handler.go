// Package handlers HTTP请求处理器
// 处理 Token 使用记录相关的 HTTP 请求
package handlers

import (
	"net/http"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type TokenUsageHandler struct {
	tokenUsageService *services.TokenUsageService
}

// NewTokenUsageHandler 创建 Token 使用记录处理器实例
func NewTokenUsageHandler() *TokenUsageHandler {
	return &TokenUsageHandler{
		tokenUsageService: services.NewTokenUsageService(),
	}
}

// RecordUsage 记录 Token 使用
func (h *TokenUsageHandler) RecordUsage(c *gin.Context) {
	var req services.RecordUsageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	if err := h.tokenUsageService.RecordUsage(&req); err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, nil)
}

// ListUsageLogs 获取使用记录列表
func (h *TokenUsageHandler) ListUsageLogs(c *gin.Context) {
	var req services.ListUsageLogsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.tokenUsageService.ListUsageLogs(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}
