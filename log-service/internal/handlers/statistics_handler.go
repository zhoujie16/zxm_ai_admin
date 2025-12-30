// Package handlers 统计接口处理器
// 处理用户统计数据相关的 HTTP 请求
package handlers

import (
	"zxm_ai_admin/log-service/internal/services"

	"github.com/gin-gonic/gin"
)

// StatisticsHandler 统计处理器
type StatisticsHandler struct {
	statisticsService *services.StatisticsService
}

// NewStatisticsHandler 创建统计处理器实例
func NewStatisticsHandler() *StatisticsHandler {
	return &StatisticsHandler{
		statisticsService: services.NewStatisticsService(),
	}
}

// GetUserStatistics 获取用户统计数据
// @Summary 获取用户统计数据
// @Description 根据 authorization 字段统计用户请求的各项数据指标
// @Tags 统计
// @Accept json
// @Produce json
// @Param authorization query string true "用户唯一标识（authorization）"
// @Param start_time query string false "开始时间，格式：2006-01-02 15:04:05"
// @Param end_time query string false "结束时间，格式：2006-01-02 15:04:05"
// @Success 200 {object} services.UserStatisticsResponse
// @Router /api/request-logs/statistics [get]
func (h *StatisticsHandler) GetUserStatistics(c *gin.Context) {
	var req services.GetUserStatisticsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.statisticsService.GetUserStatistics(&req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, response)
}

// GetAuthorizationRanking 获取 authorization 使用次数排行
// @Summary 获取 authorization 使用次数排行
// @Description 统计 authorization 使用次数并按次数降序排列
// @Tags 统计
// @Accept json
// @Produce json
// @Param start_time query string false "开始时间，格式：2006-01-02 15:04:05"
// @Param end_time query string false "结束时间，格式：2006-01-02 15:04:05"
// @Param page query int false "页码，默认 1"
// @Param page_size query int false "每页数量，默认 20，最大 100"
// @Success 200 {object} services.AuthorizationRankingResponse
// @Router /api/request-logs/ranking [get]
func (h *StatisticsHandler) GetAuthorizationRanking(c *gin.Context) {
	var req services.GetAuthorizationRankingRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.statisticsService.GetAuthorizationRanking(&req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, response)
}
