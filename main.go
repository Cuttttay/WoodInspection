package main

import (
	"WoodInspection/api"
	"WoodInspection/internal/middleware"
	"WoodInspection/internal/product/auth/app"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 1. 创建依赖注入容器（自动完成所有初始化，包括 RSA 公钥）
	container, err := app.NewContainer("./config")
	if err != nil {
		log.Fatalf("初始化容器失败: %v", err)
	}
	defer func() {
		if err := container.Close(); err != nil {
			log.Printf("关闭资源失败: %v", err)
		}
	}()

	// 3. 创建 Gin 引擎
	r := gin.Default()

	// 3.1 添加 CORS 中间件（必须在路由之前）
	r.Use(middleware.CORSMiddleware())

	// 4. 设置路由
	api.SetupRoutes(r, container)

	// 5. 启动服务器
	port := container.Config.App.Port
	addr := ":" + strconv.Itoa(port)

	log.Printf("服务器启动在端口: %d\n", port)

	// 6. 优雅关闭
	go func() {
		if err := r.Run(addr); err != nil {
			log.Fatalf("服务器启动失败: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("\n正在关闭服务器...")
}

func main1() {
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	fmt.Println(string(hash))
}
