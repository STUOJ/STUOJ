package routes

import (
	"STUOJ/server/handler"

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
		problemAdminRoute.POST("/problem", handler.ProblemAdd)
		problemAdminRoute.PUT("/problem", handler.ProblemModify)
		problemAdminRoute.DELETE("/problem/:id", handler.ProblemRemove)
		problemAdminRoute.POST("/problem/tag", handler.ProblemAddTag)
		problemAdminRoute.DELETE("/problem/tag", handler.ProblemRemoveTag)
		problemAdminRoute.POST("/problem/fps", handler.ProblemParseFromFps)
		problemAdminRoute.GET("/history/problem/:id", handler.HistoryListOfProblem)
	}
}
