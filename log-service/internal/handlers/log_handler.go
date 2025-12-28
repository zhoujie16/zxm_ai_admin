// Package handlers HTTP请求处理器
// 处理日志相关的 HTTP 请求
package handlers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
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

// CreateLog 创建日志记录（proxy 调用）
// @Summary 创建日志记录
// @Description 记录一次 Token 请求的使用情况
// @Tags 日志
// @Accept json
// @Produce json
// @Param request body services.CreateLogRequest true "日志信息"
// @Success 200 {object} models.TokenUsageLog
// @Router /api/logs [post]
func (h *LogHandler) CreateLog(c *gin.Context) {
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

// UploadLog 上传日志文件（proxy 调用）
// @Summary 上传日志文件
// @Description 接收 proxy 上传的日志文件，解析后保存到数据库
// @Tags 日志
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "日志文件"
// @Param type formData string true "日志类型" Enums(request, system)
// @Success 200 {object} map[string]interface{}
// @Router /api/logs/upload [post]
func (h *LogHandler) UploadLog(c *gin.Context) {
	// 获取上传的文件
	fileHeader, err := c.FormFile("file")
	if err != nil {
		BadRequest(c, "获取文件失败: "+err.Error())
		return
	}

	// 获取日志类型
	logType := c.PostForm("type")
	if logType != "request" && logType != "system" {
		BadRequest(c, "日志类型必须是 request 或 system")
		return
	}

	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		InternalServerError(c, "打开文件失败: "+err.Error())
		return
	}
	defer file.Close()

	// 逐行解析并保存
	successCount := 0
	failCount := 0
	scanner := bufio.NewScanner(file)
	// 增加缓冲区大小，默认 64KB 可能不够处理长日志行
	const maxScanTokenSize = 1024 * 1024 // 1MB
	buf := make([]byte, 0, maxScanTokenSize)
	scanner.Buffer(buf, maxScanTokenSize)
	batchSize := 100

	if logType == "request" {
		// 处理请求日志
		var batch []services.CreateLogRequest
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var entry map[string]interface{}
			if err := json.Unmarshal([]byte(line), &entry); err != nil {
				failCount++
				continue
			}

			req := h.parseRequestLogEntry(entry)
			if req != nil {
				batch = append(batch, *req)
				if len(batch) >= batchSize {
					if h.logService.BatchCreateLogs(batch) == len(batch) {
						successCount += len(batch)
					} else {
						failCount += len(batch)
					}
					batch = nil
				}
			} else {
				failCount++
			}
		}

		// 处理剩余的批次
		if len(batch) > 0 {
			if h.logService.BatchCreateLogs(batch) == len(batch) {
				successCount += len(batch)
			} else {
				failCount += len(batch)
			}
		}
	} else {
		// 处理系统日志
		var batch []services.CreateSystemLogRequest
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}

			var entry map[string]interface{}
			if err := json.Unmarshal([]byte(line), &entry); err != nil {
				failCount++
				continue
			}

			req := h.parseSystemLogEntry(entry)
			if req != nil {
				batch = append(batch, *req)
				if len(batch) >= batchSize {
					if count, err := h.systemLogService.BatchCreateSystemLogs(batch); err == nil {
						successCount += count
					} else {
						failCount += len(batch)
					}
					batch = nil
				}
			} else {
				failCount++
			}
		}

		// 处理剩余的批次
		if len(batch) > 0 {
			if count, err := h.systemLogService.BatchCreateSystemLogs(batch); err == nil {
				successCount += count
			} else {
				failCount += len(batch)
			}
		}
	}

	// 检查 scanner 错误，但不影响已处理的结果
	scanErr := scanner.Err()
	if scanErr != nil {
		// 记录错误但仍然返回成功/失败计数
		failCount++
	}

	Success(c, gin.H{
		"success_count": successCount,
		"fail_count":    failCount,
		"scan_error":    scanErr != nil,
	})
}

// parseRequestLogEntry 解析请求日志条目
func (h *LogHandler) parseRequestLogEntry(entry map[string]interface{}) *services.CreateLogRequest {
	// 解析时间
	timeStr, _ := entry["time"].(string)
	logTime, err := time.Parse("2006-01-02 15:04:05.000", timeStr)
	if err != nil {
		// 尝试其他格式
		logTime, _ = time.Parse(time.RFC3339, timeStr)
	}
	if logTime.IsZero() {
		logTime = time.Now()
	}

	// 解析 level
	level, _ := entry["level"].(string)
	// slog 的 level 可能是数字格式，需要转换
	if level == "" {
		if lvl, ok := entry["level"].(float64); ok {
			switch int(lvl) {
			case -4:
				level = "DEBUG"
			case 0:
				level = "INFO"
			case 4:
				level = "WARN"
			case 8:
				level = "ERROR"
			default:
				level = "INFO"
			}
		}
	}

	msg, _ := entry["msg"].(string)

	req := &services.CreateLogRequest{
		Time:  logTime,
		Level: level,
		Msg:   msg,
	}

	// 提取常用字段
	if v, ok := entry["request_id"].(string); ok {
		req.RequestID = v
	}
	if v, ok := entry["method"].(string); ok {
		req.Method = v
	}
	if v, ok := entry["path"].(string); ok {
		req.Path = v
	}
	if v, ok := entry["query"].(string); ok {
		req.Query = v
	}
	if v, ok := entry["remote_addr"].(string); ok {
		req.RemoteAddr = v
	}
	if v, ok := entry["user_agent"].(string); ok {
		req.UserAgent = v
	}
	if v, ok := entry["x_forwarded_for"].(string); ok {
		req.XForwardedFor = v
	}
	if v, ok := entry["authorization"].(string); ok {
		req.Authorization = v
	}
	if v, ok := entry["request_body"].(string); ok {
		req.RequestBody = v
	}
	if v, ok := entry["status"].(float64); ok {
		req.Status = int(v)
	}
	if v, ok := entry["latency_ms"].(float64); ok {
		req.LatencyMs = int64(v)
	}
	if v, ok := entry["request_size_bytes"].(float64); ok {
		req.RequestSizeBytes = int(v)
	}
	if v, ok := entry["response_size_bytes"].(float64); ok {
		req.ResponseSizeBytes = int(v)
	}

	// 解析 request_headers
	if v, ok := entry["request_headers"].(map[string]interface{}); ok {
		req.RequestHeaders = make(map[string]string, len(v))
		for key, val := range v {
			if strVal, ok := val.(string); ok {
				req.RequestHeaders[key] = strVal
			}
		}
	} else {
		req.RequestHeaders = make(map[string]string)
	}

	// 解析 response_headers
	if v, ok := entry["response_headers"].(map[string]interface{}); ok {
		req.ResponseHeaders = make(map[string]string, len(v))
		for key, val := range v {
			if strVal, ok := val.(string); ok {
				req.ResponseHeaders[key] = strVal
			}
		}
	} else {
		req.ResponseHeaders = make(map[string]string)
	}

	// 请求日志必须有 request_id，否则跳过
	if req.RequestID == "" {
		return nil
	}

	return req
}

// parseSystemLogEntry 解析系统日志条目
func (h *LogHandler) parseSystemLogEntry(entry map[string]interface{}) *services.CreateSystemLogRequest {
	// 解析时间
	timeStr, _ := entry["time"].(string)
	logTime, err := time.Parse("2006-01-02 15:04:05.000", timeStr)
	if err != nil {
		logTime, _ = time.Parse(time.RFC3339, timeStr)
	}
	if logTime.IsZero() {
		logTime = time.Now()
	}

	// 解析 level
	level, _ := entry["level"].(string)
	if level == "" {
		if lvl, ok := entry["level"].(float64); ok {
			switch int(lvl) {
			case -4:
				level = "DEBUG"
			case 0:
				level = "INFO"
			case 4:
				level = "WARN"
			case 8:
				level = "ERROR"
			default:
				level = "INFO"
			}
		}
	}

	msg, _ := entry["msg"].(string)

	return &services.CreateSystemLogRequest{
		Time:  logTime,
		Level: level,
		Msg:   msg,
	}
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
