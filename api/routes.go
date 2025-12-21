package api

import (
	"WoodInspection/internal/middleware"
	"WoodInspection/internal/product/auth/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, container *app.Container) {
	var api = r.Group("/api")
	{
		api.POST("/login", container.UserController.UserLogin)
		api.POST("/dectect")
	}
	auth := api.Group("/user")
	auth.Use(middleware.AuthMiddleWare([]byte(container.Config.JWT.Secret)))
	{
		// 获取当前用户信息
		auth.GET("/info", container.UserController.GetUserInfo)

	}
	// 管理员接口（需要 JWT 认证）
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleWare([]byte(container.Config.JWT.Secret)))
	{

	}
}
