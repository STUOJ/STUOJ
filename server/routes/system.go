package routes

import (
	"STUOJ/server/handler"
	handler_admin "STUOJ/server/handler-admin"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSystemRoute(ginServer *gin.Engine) {
	rootPrivateRoute := ginServer.Group("/system")
	{
		// 使用中间件
		rootPrivateRoute.Use(middlewares.TokenAuthRoot())

		rootPrivateRoute.PUT("/user/role", handler_admin.AdminUserModifyRole)

		rootPrivateRoute.GET("/config", handler.ConfigList)
	}
}
