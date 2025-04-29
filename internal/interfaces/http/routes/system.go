package routes

import (
	"STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSystemRoute(ginServer *gin.Engine) {
	rootPrivateRoute := ginServer.Group("/system")
	{
		// 使用中间件
		rootPrivateRoute.Use(middlewares.TokenAuthRoot())

		rootPrivateRoute.GET("/config", handler.ConfigList)
	}
}
