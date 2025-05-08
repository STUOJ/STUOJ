package routes

import (
	"STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

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

	userUserRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userUserRoute.Use(middlewares.TokenAuthUser())

		userUserRoute.GET("/current", handler.UserCurrentId)
		userUserRoute.PUT("/modify", handler.UserModify)
		userUserRoute.POST("/avatar", handler.ModifyUserAvatar)
	}

	userAdminRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userAdminRoute.Use(middlewares.TokenAuthAdmin())

		userAdminRoute.GET("/", handler.UserList)
		userAdminRoute.POST("/", handler.UserAdd)
		userAdminRoute.PUT("/role", handler.UserModifyRole)
		userAdminRoute.GET("/statistics", handler.UserStatistics)

	}

	userRootRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userRootRoute.Use(middlewares.TokenAuthRoot())
		userRootRoute.DELETE("/:id", handler.UserRemove)
	}
}
