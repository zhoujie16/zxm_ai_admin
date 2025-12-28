// Package middleware HTTP中间件
// 错误恢复中间件，捕获panic并返回友好的错误响应
package middleware

import (
	"zxm_ai_admin/server/internal/logger"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 错误恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic recovered",
					"error", err,
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
				)
				utils.InternalServerError(c, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
