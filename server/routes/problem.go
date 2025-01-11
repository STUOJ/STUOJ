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

		problemPublicRoute.GET("/tag", handler.TagList)
	}
	problemAdminRoute := ginServer.Group("/problem")
	{
		// 使用中间件
		problemAdminRoute.Use(middlewares.TokenAuthAdmin())

		problemAdminRoute.POST("/", handler.ProblemAdd)
		problemAdminRoute.PUT("/", handler.ProblemModify)
		problemAdminRoute.DELETE("/:id", handler.ProblemRemove)
		problemAdminRoute.POST("/tag", handler.ProblemAddTag)
		problemAdminRoute.DELETE("/tag", handler.ProblemRemoveTag)
		problemAdminRoute.POST("/fps", handler.ProblemParseFromFps)
		problemAdminRoute.GET("/history/:id", handler.HistoryListOfProblem)
	}
}
