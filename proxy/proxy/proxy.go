package proxy

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"goTest/logger"
)

// Proxy 反向代理
type Proxy struct {
	targetURL         string
	overrideAuthToken string
	proxy             *httputil.ReverseProxy
}

// New 创建反向代理
func New(targetURL string, overrideAuthToken string) (*Proxy, error) {
	target, err := url.Parse(targetURL)
	if err != nil {
		return nil, err
	}

	p := httputil.NewSingleHostReverseProxy(target)

	// 自定义 Director
	oldDirector := p.Director
	p.Director = func(req *http.Request) {
		oldDirector(req)
		req.Host = req.URL.Host
		if overrideAuthToken != "" {
			req.Header.Set("Authorization", overrideAuthToken)
		}
	}

	// 自定义 Transport
	p.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	// 错误处理
	p.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		requestID := logger.RequestIDFromContext(r.Context())
		logger.Error("代理请求失败",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"error", err.Error(),
		)
		http.Error(w, "代理请求失败", http.StatusBadGateway)
	}

	return &Proxy{
		targetURL:         targetURL,
		overrideAuthToken: overrideAuthToken,
		proxy:             p,
	}, nil
}

// Handler 返回代理处理器
func (p *Proxy) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()

		// 获取原始 Authorization
		originalAuth := r.Header.Get("Authorization")

		// 读取请求体（需要先保存以便后续代理使用）
		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = io.ReadAll(r.Body)
			r.Body.Close()
			// 重新设置 Body 以便代理使用
			r.Body = io.NopCloser(bytes.NewReader(requestBody))
		}

		// 使用响应包装器
		wrapped := &ResponseWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		// 执行代理
		p.proxy.ServeHTTP(wrapped, r)

		// 记录日志
		p.logRequest(ctx, r, wrapped, originalAuth, string(requestBody), start)
	})
}

func (p *Proxy) logRequest(ctx context.Context, r *http.Request, wrapped *ResponseWrapper, originalAuth, requestBody string, start time.Time) {
	latency := time.Since(start)
	statusCode := wrapped.StatusCode

	// 根据状态码决定日志级别
	level := slog.LevelInfo
	msg := "proxy_request"
	if statusCode >= 400 && statusCode < 500 {
		level = slog.LevelWarn
		msg = "proxy_request_client_error"
	} else if statusCode >= 500 {
		level = slog.LevelError
		msg = "proxy_request_server_error"
	}

	logger.LogRequest(ctx, level, msg,
		"request_id", logger.RequestIDFromContext(ctx),
		"method", r.Method,
		"path", r.URL.Path,
		"query", r.URL.RawQuery,
		"remote_addr", r.RemoteAddr,
		"request_headers", headersToMap(r.Header),
		"request_body", requestBody,
		"authorization", originalAuth,
		"status", statusCode,
		"response_headers", headersToMap(wrapped.Headers),
		"response_body", wrapped.ResponseBody.String(),
		"latency_ms", latency.Milliseconds(),
		"response_size_bytes", wrapped.ResponseBody.Len(),
	)
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
