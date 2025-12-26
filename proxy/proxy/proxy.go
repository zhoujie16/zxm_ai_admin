package proxy

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"proxy/cache"
	"proxy/logger"
)

// Proxy 代理管理器
type Proxy struct {
	proxyPool   map[string]*httputil.ReverseProxy // 按目标 URL 缓存的代理池
	proxyPoolMu sync.RWMutex
	tokenCache  *cache.TokenCache // token 缓存
}

// New 创建代理
func New(tokenCache *cache.TokenCache) *Proxy {
	return &Proxy{
		proxyPool:  make(map[string]*httputil.ReverseProxy),
		tokenCache: tokenCache,
	}
}

// Handler 返回代理处理器
func (p *Proxy) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()

		originalAuth := r.Header.Get("Authorization")

		// 读取请求体
		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = io.ReadAll(r.Body)
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewReader(requestBody))
		}

		// 使用响应包装器
		wrapped := &ResponseWrapper{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		// 执行代理
		p.serveHTTP(wrapped, r)

		// 记录日志
		p.logRequest(ctx, r, wrapped, originalAuth, string(requestBody), start)
	})
}

// serveHTTP 执行代理逻辑
func (p *Proxy) serveHTTP(wrapped *ResponseWrapper, r *http.Request) {
	// 检查缓存是否就绪
	if !p.tokenCache.Ready() {
		requestID := logger.RequestIDFromContext(r.Context())
		logger.Warn("缓存未就绪，拒绝请求",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)
		http.Error(wrapped, "Service Unavailable: cache not ready", http.StatusServiceUnavailable)
		return
	}

	authHeader := r.Header.Get("Authorization")

	// 查找 token 配置
	model, exists := p.tokenCache.Lookup(authHeader)
	if !exists {
		requestID := logger.RequestIDFromContext(r.Context())
		logger.Warn("token 不在缓存中，拒绝请求",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"authorization", authHeader,
		)
		http.Error(wrapped, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// 获取或创建目标代理
	targetProxy := p.getProxy(model.AIModelAPIURL)
	if targetProxy == nil {
		requestID := logger.RequestIDFromContext(r.Context())
		logger.Error("无法创建目标代理",
			"request_id", requestID,
			"method", r.Method,
			"path", r.URL.Path,
			"target_url", model.AIModelAPIURL,
		)
		http.Error(wrapped, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 使用模型的 API Key 替换 Authorization
	p.proxyWithAPIKey(targetProxy, wrapped, r, "Bearer "+model.AIModelAPIKey)
}

// getProxy 获取或创建指定目标的代理
func (p *Proxy) getProxy(targetURL string) *httputil.ReverseProxy {
	p.proxyPoolMu.RLock()
	proxy, exists := p.proxyPool[targetURL]
	p.proxyPoolMu.RUnlock()

	if exists {
		return proxy
	}

	// 创建新代理
	p.proxyPoolMu.Lock()
	defer p.proxyPoolMu.Unlock()

	// 双重检查
	if proxy, exists := p.proxyPool[targetURL]; exists {
		return proxy
	}

	newProxy, err := createReverseProxy(targetURL)
	if err != nil {
		logger.Error("创建代理失败", "target", targetURL, "error", err)
		return nil
	}

	p.proxyPool[targetURL] = newProxy
	return newProxy
}

// proxyWithAPIKey 使用指定的 API Key 执行代理
func (p *Proxy) proxyWithAPIKey(proxy *httputil.ReverseProxy, wrapped *ResponseWrapper, r *http.Request, apiKey string) {
	// 保存原始 Authorization
	originalAuth := r.Header.Get("Authorization")
	// 设置新的 Authorization
	r.Header.Set("Authorization", apiKey)

	// 执行代理
	proxy.ServeHTTP(wrapped, r)

	// 恢复原始 Authorization（用于日志记录）
	r.Header.Set("Authorization", originalAuth)
}

// createReverseProxy 创建反向代理
func createReverseProxy(targetURL string) (*httputil.ReverseProxy, error) {
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

	return p, nil
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
		"user_agent", r.Header.Get("User-Agent"),
		"x_forwarded_for", r.Header.Get("X-Forwarded-For"),
		"request_headers", headersToMap(r.Header),
		"authorization", originalAuth,
		"request_body", requestBody,
		"status", statusCode,
		"response_headers", headersToMap(wrapped.Headers),
		"latency_ms", latency.Milliseconds(),
		"request_size_bytes", len(requestBody),
		"response_size_bytes", wrapped.ResponseSize,
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
