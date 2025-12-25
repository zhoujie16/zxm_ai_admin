// Package middleware HTTP中间件
// JWT认证中间件，验证请求中的Bearer Token并解析用户信息到上下文
package middleware

import (
	"strings"

	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "未提供认证token")
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c, "token格式错误")
			c.Abort()
			return
		}

		token := parts[1]

		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.Unauthorized(c, "无效的token")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

