// Package main 应用程序入口
// 负责初始化配置、数据库、日志等组件，启动HTTP服务器并处理优雅关闭
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"zxm_ai_admin/server/internal/config"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/handlers"
	"zxm_ai_admin/server/internal/logger"
	"zxm_ai_admin/server/internal/middleware"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	configPath := "configs/config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if err := config.Load(configPath); err != nil {
		slog.Default().Error("加载配置失败", "error", err)
		os.Exit(1)
	}

	cfg := config.GetConfig()

	// 初始化日志
	logLevel := parseLogLevel(cfg.Log.Level)
	logDir := cfg.Log.Dir
	if err := logger.System.Init(logDir, logLevel); err != nil {
		slog.Default().Error("系统日志初始化失败", "error", err)
		os.Exit(1)
	}

	// 初始化JWT
	utils.InitJWT()

	// 初始化数据库
	if err := database.Init(); err != nil {
		logger.Error("初始化数据库失败", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	r := gin.New()

	// 注册中间件
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.CORSMiddleware())

	// 注册路由
	setupRoutes(r)

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	// 启动服务器（在goroutine中）
	go func() {
		logger.Info("服务器启动", "port", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("服务器启动失败", "error", err)
			os.Exit(1)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("正在关闭服务器...")

	// 优雅关闭，等待5秒
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("服务器强制关闭", "error", err)
		os.Exit(1)
	}

	logger.Info("服务器已关闭")
}

// parseLogLevel 解析日志级别
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// setupRoutes 设置路由
func setupRoutes(r *gin.Engine) {
	// 健康检查
	healthHandler := handlers.NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	// API路由组
	api := r.Group("/api")
	{
		// 认证相关
		authHandler := handlers.NewAuthHandler()
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.GET("/me", middleware.AuthMiddleware(), authHandler.GetMe)
		}

		// 代理服务相关（需要认证）
		proxyServiceHandler := handlers.NewProxyServiceHandler()
		proxyServices := api.Group("/proxy-services")
		proxyServices.Use(middleware.AuthMiddleware())
		{
			proxyServices.POST("", proxyServiceHandler.CreateProxyService)
			proxyServices.GET("", proxyServiceHandler.ListProxyServices)
			proxyServices.GET("/:id", proxyServiceHandler.GetProxyService)
			proxyServices.PUT("/:id", proxyServiceHandler.UpdateProxyService)
			proxyServices.DELETE("/:id", proxyServiceHandler.DeleteProxyService)
		}

		// AI 模型相关（需要认证）
		aiModelHandler := handlers.NewAIModelHandler()
		aiModels := api.Group("/ai-models")
		aiModels.Use(middleware.AuthMiddleware())
		{
			aiModels.POST("", aiModelHandler.CreateAIModel)
			aiModels.GET("", aiModelHandler.ListAIModels)
			aiModels.GET("/:id", aiModelHandler.GetAIModel)
			aiModels.PUT("/:id", aiModelHandler.UpdateAIModel)
			aiModels.DELETE("/:id", aiModelHandler.DeleteAIModel)
		}

		// Token 相关（需要认证）
		tokenHandler := handlers.NewTokenHandler()
		tokens := api.Group("/tokens")
		tokens.Use(middleware.AuthMiddleware())
		{
			tokens.POST("", tokenHandler.CreateToken)
			tokens.GET("", tokenHandler.ListTokens)
			tokens.GET("/recycle", tokenHandler.ListRecycledTokens)
			tokens.GET("/:id", tokenHandler.GetToken)
			tokens.PUT("/:id", tokenHandler.UpdateToken)
			tokens.DELETE("/:id", tokenHandler.DeleteToken)
			tokens.POST("/:id/restore", tokenHandler.RestoreToken)
			tokens.DELETE("/:id/destroy", tokenHandler.DestroyToken)
		}

		// Token 与模型关联接口（proxy 使用系统认证令牌调用）
		api.GET("/tokens/with-model", middleware.SystemAuthMiddleware(), tokenHandler.ListAllTokensWithModel)

	}
}

