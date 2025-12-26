package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"goTest/cache"
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

	// 创建 token 缓存
	tokenCache := cache.New(cfg.ServerAPIURL, cfg.ServerAPIToken)
	cacheDone := make(chan struct{})

	// 启动缓存同步（异步）
	go tokenCache.StartSync(cfg.SyncInterval, cacheDone)

	logger.Info("token 动态路由已启用",
		"server_api", cfg.ServerAPIURL,
		"sync_interval_minutes", cfg.SyncInterval,
	)

	// 创建反向代理
	p := proxy.New(tokenCache)

	// 构建处理器链：RequestID -> Proxy
	handler := p.Handler()
	handler = middleware.RequestID(handler)

	// 启动服务
	logger.Info("代理服务器启动",
		"listen_addr", cfg.ListenAddr,
		"log_level", cfg.LogLevel,
	)

	// 监听系统信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 启动 HTTP 服务器
	go func() {
		if err := listenAndServe(cfg.ListenAddr, handler); err != nil {
			logger.Error("服务器启动失败", "error", err)
			close(cacheDone)
			os.Exit(1)
		}
	}()

	// 等待退出信号
	<-sigChan
	logger.Info("服务器正在关闭...")

	// 停止缓存同步
	close(cacheDone)

	logger.Info("服务器已关闭")
}

// listenAndServe 封装 http.ListenAndServe 以便测试
var listenAndServe = func(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}
