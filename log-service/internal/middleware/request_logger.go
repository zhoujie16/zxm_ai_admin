// Package middleware 中间件
// HTTP 请求日志中间件，记录所有 HTTP 请求的详细信息
package middleware

import (
	"bytes"
	"io"
	"log/slog"
	"time"

	"zxm_ai_admin/log-service/internal/logger"

	"github.com/gin-gonic/gin"
)

// bodyLogWriter 用于捕获响应体
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// RequestLogger HTTP 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 读取请求体
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				requestBody = string(bodyBytes)
				// 限制请求体长度，避免日志过大
				if len(requestBody) > 10000 {
					requestBody = requestBody[:10000] + "...(truncated)"
				}
			}
		}

		// 捕获响应体
		blw := &bodyLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 计算延迟
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// 根据状态码决定日志级别
		level := slog.LevelInfo
		msg := "http_request"
		if statusCode >= 400 && statusCode < 500 {
			level = slog.LevelWarn
			msg = "http_request_client_error"
		} else if statusCode >= 500 {
			level = slog.LevelError
			msg = "http_request_server_error"
		}

		// 记录日志
		args := []any{
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"remote_addr", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"status", statusCode,
			"latency_ms", latency.Milliseconds(),
			"request_size_bytes", len(requestBody),
		}

		// 添加请求体（仅对非 GET 请求且长度合理时）
		if c.Request.Method != "GET" && len(requestBody) > 0 && len(requestBody) <= 1000 {
			args = append(args, "request_body", requestBody)
		}

		// 添加响应体（仅对错误响应且长度合理时）
		responseBody := blw.body.String()
		if statusCode >= 400 && len(responseBody) > 0 && len(responseBody) <= 1000 {
			args = append(args, "response_body", responseBody)
		}

		logger.System.Log(c.Request.Context(), level, msg, args...)
	}
}

