package main

import (
	"net/http"
	"os"

	"goTest/config"
	"goTest/logger"
	"goTest/middleware"
	"goTest/proxy"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志
	logger.Init(config.ParseLogLevel(cfg.LogLevel))

	// 创建反向代理
	p, err := proxy.New(cfg.TargetURL, cfg.OverrideAuthToken)
	if err != nil {
		logger.Error("无效的目标 URL", "error", err)
		os.Exit(1)
	}

	// 构建处理器链：RequestID -> Auth -> Proxy
	handler := p.Handler()
	if len(cfg.RequiredAuthTokens) > 0 {
		handler = middleware.Auth(cfg.RequiredAuthTokens)(handler)
	}
	handler = middleware.RequestID(handler)

	// 启动服务
	logger.Info("代理服务器启动",
		"listen_addr", cfg.ListenAddr,
		"target", cfg.TargetURL,
		"log_level", cfg.LogLevel,
		"auth_enabled", len(cfg.RequiredAuthTokens) > 0,
	)

	if err := listenAndServe(cfg.ListenAddr, handler); err != nil {
		logger.Error("服务器启动失败", "error", err)
		os.Exit(1)
	}
}

// listenAndServe 封装 http.ListenAndServe 以便测试
var listenAndServe = func(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
