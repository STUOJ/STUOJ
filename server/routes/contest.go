package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitContestRoute(ginServer *gin.Engine) {
	contestPublicRoute := ginServer.Group("/contest")
	{
		contestPublicRoute.GET("/:id", handler.ContestInfo)
		contestPublicRoute.GET("/", handler.ContestList)
	}

	contestAdminRoute := ginServer.Group("/contest")
	{
		// 使用中间件
		contestAdminRoute.Use(middlewares.TokenAuthAdmin())

		contestAdminRoute.POST("/", handler.ContestAdd)
		contestAdminRoute.PUT("/", handler.ContestModify)
		contestAdminRoute.DELETE("/:id", handler.ContestRemove)
	}
}
