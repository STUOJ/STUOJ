package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitProblemRoute(ginServer *gin.Engine) {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", handler.ProblemList)
		problemPublicRoute.GET("/:id", handler.ProblemInfo)
	}

	problemEditorRoute := ginServer.Group("/problem")
	{
		// 使用中间件
		problemEditorRoute.Use(middlewares.TokenAuthAdmin())

		problemEditorRoute.POST("/", handler.ProblemAdd)
		problemEditorRoute.PUT("/", handler.ProblemModify)
		problemEditorRoute.DELETE("/:id", handler.ProblemRemove)
		problemEditorRoute.POST("/tag", handler.ProblemAddTag)
		problemEditorRoute.DELETE("/tag", handler.ProblemRemoveTag)
		problemEditorRoute.POST("/fps", handler.ProblemParseFromFps)
		problemEditorRoute.GET("/history/:id", handler.HistoryListOfProblem)
	}
}
