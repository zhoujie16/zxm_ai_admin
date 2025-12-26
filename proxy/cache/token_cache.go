package cache

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"goTest/logger"
)

// TokenModel token 与模型 的绑定关系
type TokenModel struct {
	TokenID          int    `json:"token_id"`
	Token            string `json:"token"`
	TokenStatus      int    `json:"token_status"`
	AIModelID        int    `json:"ai_model_id"`
	AIModelName      string `json:"ai_model_name"`
	AIModelAPIURL    string `json:"ai_model_api_url"`
	AIModelAPIKey    string `json:"ai_model_api_key"`
	AIModelStatus    int    `json:"ai_model_status"`
}

// APIResponse 后端 API 响应结构
type APIResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    []TokenModel `json:"data"`
}

// TokenCache token 缓存
type TokenCache struct {
	mu        sync.RWMutex
	cache     map[string]*TokenModel // key: token (sk-xxx)
	ready     bool                   // 缓存是否已就绪
	client    *http.Client
	serverURL string
	apiToken  string
}

// New 创建 token 缓存
func New(serverURL, apiToken string) *TokenCache {
	return &TokenCache{
		cache:     make(map[string]*TokenModel),
		ready:     false,
		client:    &http.Client{Timeout: 30 * time.Second},
		serverURL: serverURL,
		apiToken:  apiToken,
	}
}

// Ready 缓存是否已就绪
func (c *TokenCache) Ready() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.ready
}

// Lookup 查找 token 对应的模型配置
func (c *TokenCache) Lookup(authHeader string) (*TokenModel, bool) {
	// 提取 Bearer token
	token := extractBearerToken(authHeader)
	if token == "" {
		return nil, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	model, exists := c.cache[token]
	return model, exists
}

// extractBearerToken 从 Authorization 头中提取 token
func extractBearerToken(authHeader string) string {
	if authHeader == "" {
		return ""
	}
	// 支持 "Bearer sk-xxx" 或直接 "sk-xxx" 格式
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return strings.TrimSpace(parts[1])
	}
	return strings.TrimSpace(authHeader)
}

// Sync 同步缓存
func (c *TokenCache) Sync() error {
	url := c.serverURL + "/api/tokens/with-model"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	if c.apiToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiToken)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &APIError{StatusCode: resp.StatusCode, Message: "API 返回状态码: " + http.StatusText(resp.StatusCode)}
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return err
	}

	if apiResp.Code != 0 {
		return &APIError{StatusCode: resp.StatusCode, Message: apiResp.Message}
	}

	c.updateCache(apiResp.Data)
	logger.Info("token 缓存同步成功", "count", len(apiResp.Data))

	return nil
}

// updateCache 更新缓存
func (c *TokenCache) updateCache(items []TokenModel) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 清空旧缓存
	c.cache = make(map[string]*TokenModel)

	// 填充新缓存
	for i := range items {
		// 只缓存状态启用的 token 和模型
		if items[i].TokenStatus == 1 && items[i].AIModelStatus == 1 {
			c.cache[items[i].Token] = &items[i]
		}
	}

	// 标记缓存已就绪
	c.ready = true
}

// StartSync 启动定时同步
func (c *TokenCache) StartSync(intervalMinutes int, done chan struct{}) {
	ticker := time.NewTicker(time.Duration(intervalMinutes) * time.Minute)
	defer ticker.Stop()

	// 启动时立即同步一次
	if err := c.Sync(); err != nil {
		logger.Warn("token 首次同步失败", "error", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := c.Sync(); err != nil {
				logger.Warn("token 缓存同步失败", "error", err)
			}
		case <-done:
			return
		}
	}
}

// APIError API 错误
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return e.Message
}
