// Package handlers HTTP请求处理器
// 处理 Token 使用记录相关的 HTTP 请求
package handlers

import (
	"net/http"
	"strconv"
	"zxm_ai_admin/server/internal/services"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

type TokenUsageLogHandler struct {
	tokenUsageLogService *services.TokenUsageLogService
}

// NewTokenUsageLogHandler 创建 Token 使用记录处理器实例
func NewTokenUsageLogHandler() *TokenUsageLogHandler {
	return &TokenUsageLogHandler{
		tokenUsageLogService: services.NewTokenUsageLogService(),
	}
}

// CreateTokenUsageLog 创建 Token 使用记录
// @Summary 创建 Token 使用记录
// @Description 记录一次 Token 请求的使用情况
// @Tags Token使用记录
// @Accept json
// @Produce json
// @Param request body services.CreateTokenUsageLogRequest true "使用记录信息"
// @Success 200 {object} models.TokenUsageLog
// @Router /api/token-usage-logs [post]
func (h *TokenUsageLogHandler) CreateTokenUsageLog(c *gin.Context) {
	var req services.CreateTokenUsageLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	log, err := h.tokenUsageLogService.CreateTokenUsageLog(&req)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(c, log)
}

// ListTokenUsageLogs 获取 Token 使用记录列表
// @Summary 获取 Token 使用记录列表
// @Description 分页查询 Token 使用记录，支持多种过滤条件
// @Tags Token使用记录
// @Accept json
// @Produce json
// @Param page query int false "页码，从1开始" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param request_id query string false "按 request_id 精确查询"
// @Param start_time query string false "开始时间（格式：2006-01-02 15:04:05）"
// @Param end_time query string false "结束时间（格式：2006-01-02 15:04:05）"
// @Param status query int false "按状态码过滤（-1 表示全部）"
// @Param method query string false "按 HTTP 方法过滤"
// @Param authorization query string false "按 Authorization 头模糊匹配"
// @Success 200 {object} services.ListTokenUsageLogsResponse
// @Router /api/token-usage-logs [get]
func (h *TokenUsageLogHandler) ListTokenUsageLogs(c *gin.Context) {
	var req services.ListTokenUsageLogsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.tokenUsageLogService.ListTokenUsageLogs(&req)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// GetTokenUsageLog 获取 Token 使用记录详情
// @Summary 获取 Token 使用记录详情
// @Description 根据 ID 获取单条 Token 使用记录
// @Tags Token使用记录
// @Accept json
// @Produce json
// @Param id path int true "记录ID"
// @Success 200 {object} models.TokenUsageLog
// @Router /api/token-usage-logs/{id} [get]
func (h *TokenUsageLogHandler) GetTokenUsageLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的ID")
		return
	}

	log, err := h.tokenUsageLogService.GetTokenUsageLog(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, log)
}
