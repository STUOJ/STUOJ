package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSolutionRoute(ginServer *gin.Engine) {
	solutionAdminRoute := ginServer.Group("/solution")
	{
		// 使用中间件
		solutionAdminRoute.Use(middlewares.TokenAuthAdmin())

		solutionAdminRoute.GET("/:id", handler.SolutionInfo)
		solutionAdminRoute.POST("/", handler.SolutionAdd)
		solutionAdminRoute.PUT("/", handler.SolutionModify)
		solutionAdminRoute.DELETE("/:id", handler.SolutionRemove)
	}
}
