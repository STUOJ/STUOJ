package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(ginServer *gin.Engine) {
	userPublicRoute := ginServer.Group("/user")
	{
		userPublicRoute.GET("/:id", handler.UserInfo)
		userPublicRoute.POST("/login", handler.UserLogin)
		userPublicRoute.POST("/register", handler.UserRegister)
		userPublicRoute.PUT("/password", handler.UserChangePassword)
	}
	userProtectedRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userProtectedRoute.Use(middlewares.TokenAuthUser())

		userProtectedRoute.GET("/current", handler.UserCurrentId)
		userProtectedRoute.PUT("/modify/:id", handler.UserModify)
		userProtectedRoute.POST("/avatar/:id", handler.ModifyUserAvatar)
	}
}
