// Package main 应用程序入口
// 负责初始化配置、数据库、日志等组件，启动HTTP服务器并处理优雅关闭
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"zxm_ai_admin/server/internal/config"
	"zxm_ai_admin/server/internal/database"
	"zxm_ai_admin/server/internal/handlers"
	"zxm_ai_admin/server/internal/middleware"
	"zxm_ai_admin/server/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 加载配置
	configPath := "configs/config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if err := config.Load(configPath); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	cfg := config.GetConfig()

	// 初始化日志
	logger := initLogger(cfg.Log.Level, cfg.Log.Output)
	defer logger.Sync()

	// 初始化JWT
	utils.InitJWT()

	// 初始化数据库
	if err := database.Init(); err != nil {
		logger.Fatal("初始化数据库失败", zap.Error(err))
	}
	defer database.Close()

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	r := gin.New()

	// 注册中间件
	r.Use(middleware.RecoveryMiddleware(logger))
	r.Use(middleware.LoggerMiddleware(logger))
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
		logger.Info("服务器启动", zap.Int("port", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("服务器启动失败", zap.Error(err))
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
		logger.Fatal("服务器强制关闭", zap.Error(err))
	}

	logger.Info("服务器已关闭")
}

// initLogger 初始化日志
func initLogger(level, output string) *zap.Logger {
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapLevel)
	config.OutputPaths = []string{output, "stdout"}
	config.ErrorOutputPaths = []string{output, "stderr"}
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	return logger
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
	}
}

