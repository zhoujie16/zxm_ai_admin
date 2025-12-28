// Package main 应用程序入口
// 负责初始化配置、数据库，启动HTTP服务器
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

	"zxm_ai_admin/log-service/internal/config"
	"zxm_ai_admin/log-service/internal/database"
	"zxm_ai_admin/log-service/internal/handlers"
	"zxm_ai_admin/log-service/internal/middleware"

	"github.com/gin-gonic/gin"
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

	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer database.Close()

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建Gin引擎
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())

	// 注册路由
	setupRoutes(r)

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: r,
	}

	// 启动服务器
	go func() {
		log.Printf("日志服务启动，端口: %d", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器强制关闭: %v", err)
	}

	log.Println("服务器已关闭")
}

// setupRoutes 设置路由
func setupRoutes(r *gin.Engine) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "ok",
		})
	})

	logHandler := handlers.NewLogHandler()

	// API路由组
	api := r.Group("/api")
	{
		// 写入日志（使用 System Auth Token 认证）
		api.POST("/request-logs", middleware.SystemAuthMiddleware(), logHandler.CreateRequestLog)
		api.POST("/request-logs/batch", middleware.SystemAuthMiddleware(), logHandler.BatchCreateRequestLogs)
		api.POST("/system-logs/batch", middleware.SystemAuthMiddleware(), logHandler.BatchCreateSystemLogs)

		// 查询日志（使用 JWT 认证）
		api.GET("/request-logs", middleware.AuthMiddleware(), logHandler.ListLogs)
		api.GET("/request-logs/:id", middleware.AuthMiddleware(), logHandler.GetLog)

		// 系统日志查询（使用 JWT 认证）
		api.GET("/system-logs", middleware.AuthMiddleware(), logHandler.ListSystemLogs)
		api.GET("/system-logs/:id", middleware.AuthMiddleware(), logHandler.GetSystemLog)

		// 删除日志（token 在 body 中验证）
		api.POST("/request-logs/delete", logHandler.DeleteRequestLogs)
		api.POST("/system-logs/delete", logHandler.DeleteSystemLogs)
	}
}
