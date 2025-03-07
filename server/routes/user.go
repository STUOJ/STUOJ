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

	userUserRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userUserRoute.Use(middlewares.TokenAuthUser())

		userUserRoute.GET("/current", handler.UserCurrentId)
		userUserRoute.PUT("/modify/:id", handler.UserModify)
		userUserRoute.POST("/avatar/:id", handler.ModifyUserAvatar)
	}

	userAdminRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userAdminRoute.Use(middlewares.TokenAuthAdmin())

		userAdminRoute.GET("/", handler.UserList)
		userAdminRoute.POST("/", handler.UserAdd)
		userAdminRoute.PUT("/role", handler.UserModifyRole)

	}

	userRootRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userRootRoute.Use(middlewares.TokenAuthRoot())
		userRootRoute.DELETE("/:id", handler.UserRemove)
	}
}
