package middleware

import (
	"net/http"

	"goTest/logger"
)

// Auth 认证中间件
func Auth(requiredTokens []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 如果没有配置白名单，直接放行
			if len(requiredTokens) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			originalAuth := r.Header.Get("Authorization")
			requestID := logger.RequestIDFromContext(r.Context())

			// 校验白名单
			allowed := false
			for _, token := range requiredTokens {
				if originalAuth == token {
					allowed = true
					break
				}
			}

			if !allowed {
				logger.Warn("请求被拒绝",
					"request_id", requestID,
					"method", r.Method,
					"path", r.URL.Path,
					"query", r.URL.RawQuery,
					"remote_addr", r.RemoteAddr,
					"request_headers", headersToMap(r.Header),
					"authorization", originalAuth,
					"reason", "not_in_whitelist",
				)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func headersToMap(h http.Header) map[string]string {
	if h == nil {
		return nil
	}
	result := make(map[string]string, len(h))
	for k, v := range h {
		if len(v) == 1 {
			result[k] = v[0]
		} else if len(v) > 1 {
			joined := ""
			for i, val := range v {
				if i > 0 {
					joined += ", "
				}
				joined += val
			}
			result[k] = joined
		}
	}
	return result
}
