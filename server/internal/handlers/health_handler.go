// Package handlers HTTP请求处理器
// 处理系统健康检查相关的HTTP请求
package handlers

import (
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

// NewHealthHandler 创建健康检查处理器实例
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Health 健康检查
// @Summary 健康检查
// @Description 服务健康检查接口
// @Tags 系统
// @Produce json
// @Success 200 {object} utils.Response
// @Router /health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	utils.Success(c, gin.H{
		"status": "ok",
		"message": "服务运行正常",
	})
}

