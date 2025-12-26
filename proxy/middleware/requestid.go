package middleware

import (
	"net/http"

	"proxy/logger"
)

// RequestID 为每个请求生成 requestID 的中间件
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := logger.GenerateRequestID()
		ctx := logger.ContextWithRequestID(r.Context(), requestID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
