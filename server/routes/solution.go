package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSolutionRoute(ginServer *gin.Engine) {
	solutionEditorRoute := ginServer.Group("/solution")
	{
		// 使用中间件
		solutionEditorRoute.Use(middlewares.TokenAuthAdmin())

		solutionEditorRoute.GET("/:id", handler.SolutionInfo)
		solutionEditorRoute.POST("/", handler.SolutionAdd)
		solutionEditorRoute.PUT("/", handler.SolutionModify)
		solutionEditorRoute.DELETE("/:id", handler.SolutionRemove)
	}
}
