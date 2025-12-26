package cache

import (
	"bytes"
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

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token     string `json:"token"`
		Username  string `json:"username"`
		UserInfo  struct {
			Username string `json:"username"`
		} `json:"user_info"`
	} `json:"data"`
}

// TokenCache token 缓存
type TokenCache struct {
	mu            sync.RWMutex
	cache         map[string]*TokenModel // key: token (sk-xxx)
	ready         bool                   // 缓存是否已就绪
	client        *http.Client
	serverBaseURL string
	username      string
	password      string
	jwtToken      string // 登录获取的 JWT token
}

// New 创建 token 缓存
func New(serverBaseURL, username, password string) *TokenCache {
	return &TokenCache{
		cache:         make(map[string]*TokenModel),
		ready:         false,
		client:        &http.Client{Timeout: 30 * time.Second},
		serverBaseURL: serverBaseURL,
		username:      username,
		password:      password,
	}
}

// Login 登录获取 JWT token
func (c *TokenCache) Login() error {
	url := c.serverBaseURL + "/api/auth/login"
	reqBody := LoginRequest{
		Username: c.username,
		Password: c.password,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &APIError{StatusCode: resp.StatusCode, Message: "登录返回状态码: " + http.StatusText(resp.StatusCode)}
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return err
	}

	if loginResp.Code != 0 {
		return &APIError{StatusCode: resp.StatusCode, Message: loginResp.Message}
	}

	c.mu.Lock()
	c.jwtToken = loginResp.Data.Token
	c.mu.Unlock()

	logger.Info("登录成功", "username", c.username)
	return nil
}

// getJWTToken 获取当前的 JWT token
func (c *TokenCache) getJWTToken() string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.jwtToken
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
	url := c.serverBaseURL + "/api/tokens/with-model"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	jwtToken := c.getJWTToken()
	if jwtToken == "" {
		return &APIError{StatusCode: 0, Message: "JWT token 为空，请先登录"}
	}

	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 401 表示 token 过期，重新登录后重试
	if resp.StatusCode == http.StatusUnauthorized {
		logger.Info("JWT token 过期，重新登录")
		if err := c.Login(); err != nil {
			return err
		}
		// 重试同步
		return c.Sync()
	}

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

	// 启动时先登录
	if err := c.Login(); err != nil {
		logger.Warn("登录失败", "error", err)
	}

	// 登录成功后立即同步一次
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
