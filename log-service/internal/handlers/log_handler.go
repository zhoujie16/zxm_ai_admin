// Package handlers HTTP请求处理器
// 处理日志相关的 HTTP 请求
package handlers

import (
	"net/http"
	"strconv"
	"zxm_ai_admin/log-service/internal/logger"
	"zxm_ai_admin/log-service/internal/services"

	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	logService       *services.LogService
	systemLogService *services.SystemLogService
}

// NewLogHandler 创建日志处理器实例
func NewLogHandler() *LogHandler {
	return &LogHandler{
		logService:       services.NewLogService(),
		systemLogService: services.NewSystemLogService(),
	}
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

// BadRequest 错误响应
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    400,
		"message": message,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":    401,
		"message": message,
	})
}

// NotFound 未找到响应
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": message,
	})
}

// InternalServerError 内部错误响应
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":    500,
		"message": message,
	})
}

// CreateRequestLog 创建请求日志记录（proxy 调用）
// @Summary 创建请求日志记录
// @Description 记录一次 Token 请求的使用情况
// @Tags 请求日志
// @Accept json
// @Produce json
// @Param request body services.CreateLogRequest true "日志信息"
// @Success 200 {object} models.TokenUsageLog
// @Router /api/request-logs [post]
func (h *LogHandler) CreateRequestLog(c *gin.Context) {
	var req services.CreateLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	log, err := h.logService.CreateLog(&req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, log)
}

// ListLogs 获取请求日志列表（admin 调用）
// @Summary 获取请求日志列表
// @Description 分页查询请求日志，支持多种过滤条件
// @Tags 请求日志
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param request_id query string false "按 request_id 查询"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Param status query string false "状态码（单个如 200 或多个逗号分隔如 200,401,404）"
// @Param method query string false "HTTP 方法"
// @Param authorization query string false "Authorization"
// @Success 200 {object} services.ListLogsResponse
// @Router /api/request-logs [get]
func (h *LogHandler) ListLogs(c *gin.Context) {
	var req services.ListLogsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.logService.ListLogs(&req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, response)
}

// GetLog 获取请求日志详情（admin 调用）
// @Summary 获取请求日志详情
// @Description 根据 ID 获取单条请求日志
// @Tags 请求日志
// @Accept json
// @Produce json
// @Param id path int true "记录ID"
// @Success 200 {object} models.TokenUsageLog
// @Router /api/request-logs/{id} [get]
func (h *LogHandler) GetLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		BadRequest(c, "无效的ID")
		return
	}

	log, err := h.logService.GetLog(uint(id))
	if err != nil {
		NotFound(c, err.Error())
		return
	}

	Success(c, log)
}

// BatchCreateRequestLogs 批量创建请求日志记录（proxy 调用）
// @Summary 批量创建请求日志记录
// @Description 接收多条请求日志记录的数组，批量保存到数据库
// @Tags 请求日志
// @Accept json
// @Produce json
// @Param request body []services.CreateLogRequest true "日志记录数组"
// @Success 200 {object} map[string]interface{}
// @Router /api/request-logs/batch [post]
func (h *LogHandler) BatchCreateRequestLogs(c *gin.Context) {
	var reqs []services.CreateLogRequest
	if err := c.ShouldBindJSON(&reqs); err != nil {
		logger.Error("批量创建请求日志失败：参数绑定错误", "error", err)
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	logger.Info("批量创建请求日志", "received_count", len(reqs))

	count := h.logService.BatchCreateLogs(reqs)

	logger.Info("批量创建请求日志完成", "received_count", len(reqs), "inserted_count", count)

	Success(c, gin.H{
		"count": count,
	})
}

// BatchCreateSystemLogs 批量创建系统日志记录（proxy 调用）
// @Summary 批量创建系统日志记录
// @Description 接收多条系统日志记录的数组，批量保存到数据库
// @Tags 系统日志
// @Accept json
// @Produce json
// @Param request body []services.CreateSystemLogRequest true "系统日志记录数组"
// @Success 200 {object} map[string]interface{}
// @Router /api/system-logs/batch [post]
func (h *LogHandler) BatchCreateSystemLogs(c *gin.Context) {
	var reqs []services.CreateSystemLogRequest
	if err := c.ShouldBindJSON(&reqs); err != nil {
		logger.Error("批量创建系统日志失败：参数绑定错误", "error", err)
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	logger.Info("批量创建系统日志", "received_count", len(reqs))

	count, err := h.systemLogService.BatchCreateSystemLogs(reqs)
	if err != nil {
		logger.Error("批量创建系统日志失败", "error", err, "received_count", len(reqs))
		InternalServerError(c, err.Error())
		return
	}

	logger.Info("批量创建系统日志完成", "received_count", len(reqs), "inserted_count", count)

	Success(c, gin.H{
		"count": count,
	})
}

// ListSystemLogs 获取系统日志列表（admin 调用）
// @Summary 获取系统日志列表
// @Description 分页查询系统日志，支持多种过滤条件
// @Tags 系统日志
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param level query string false "日志级别"
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Success 200 {object} services.ListSystemLogsResponse
// @Router /api/system-logs [get]
func (h *LogHandler) ListSystemLogs(c *gin.Context) {
	var req services.ListSystemLogsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.systemLogService.ListSystemLogs(&req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, response)
}

// GetSystemLog 获取系统日志详情（admin 调用）
// @Summary 获取系统日志详情
// @Description 根据 ID 获取单条系统日志
// @Tags 系统日志
// @Accept json
// @Produce json
// @Param id path int true "记录ID"
// @Success 200 {object} models.SystemLog
// @Router /api/system-logs/{id} [get]
func (h *LogHandler) GetSystemLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		BadRequest(c, "无效的ID")
		return
	}

	log, err := h.systemLogService.GetSystemLog(uint(id))
	if err != nil {
		NotFound(c, err.Error())
		return
	}

	Success(c, log)
}

// DeleteRequestLogs 按时间范围删除请求日志
// @Summary 按时间范围删除请求日志
// @Description 根据时间范围删除请求日志记录（硬删除）
// @Tags 请求日志
// @Accept json
// @Produce json
// @Param request body services.DeleteLogsByTimeRangeRequest true "删除参数"
// @Success 200 {object} services.DeleteLogsByTimeRangeResponse
// @Router /api/request-logs/delete [post]
func (h *LogHandler) DeleteRequestLogs(c *gin.Context) {
	var req services.DeleteLogsByTimeRangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.logService.DeleteLogsByTimeRange(&req)
	if err != nil {
		BadRequest(c, err.Error())
		return
	}

	Success(c, response)
}

// DeleteSystemLogs 按时间范围删除系统日志
// @Summary 按时间范围删除系统日志
// @Description 根据时间范围删除系统日志记录（硬删除）
// @Tags 系统日志
// @Accept json
// @Produce json
// @Param request body services.DeleteSystemLogsByTimeRangeRequest true "删除参数"
// @Success 200 {object} services.DeleteSystemLogsByTimeRangeResponse
// @Router /api/system-logs/delete [post]
func (h *LogHandler) DeleteSystemLogs(c *gin.Context) {
	var req services.DeleteSystemLogsByTimeRangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "参数错误: "+err.Error())
		return
	}

	response, err := h.systemLogService.DeleteSystemLogsByTimeRange(&req)
	if err != nil {
		BadRequest(c, err.Error())
		return
	}

	Success(c, response)
}
