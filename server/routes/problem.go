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
		problemEditorRoute.Use(middlewares.TokenAuthEditor())

		problemEditorRoute.POST("/", handler.ProblemAdd)
		problemEditorRoute.PUT("/", handler.ProblemModify)
		problemEditorRoute.POST("/tag", handler.ProblemAddTag)
		problemEditorRoute.DELETE("/tag", handler.ProblemRemoveTag)
		problemEditorRoute.POST("/fps", handler.ProblemParseFromFps)
		problemEditorRoute.GET("/history/:id", handler.HistoryListOfProblem)
	}
	problemAdminRoute := ginServer.Group("/problem")
	{
		// 使用中间件
		problemAdminRoute.Use(middlewares.TokenAuthAdmin())
		problemEditorRoute.DELETE("/:id", handler.ProblemRemove)
	}
}

func InitTagRoute(ginServer *gin.Engine) {
	tagPublicRoute := ginServer.Group("/tag")
	{
		tagPublicRoute.GET("/", handler.TagList)
	}

	tagEditorRoute := ginServer.Group("/tag")
	{
		// 使用中间件
		tagEditorRoute.Use(middlewares.TokenAuthEditor())
		tagEditorRoute.POST("/", handler.TagAdd)
		tagEditorRoute.PUT("/", handler.TagModify)
	}

	tagAdminRoute := ginServer.Group("/tag")
	{
		// 使用中间件
		tagAdminRoute.Use(middlewares.TokenAuthAdmin())
		tagAdminRoute.DELETE("/:id", handler.TagRemove)
	}
}

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthAdmin())

		testcaseEditorRoute.GET("/:id", handler.TestcaseInfo)
		testcaseEditorRoute.POST("/", handler.TestcaseAdd)
		testcaseEditorRoute.PUT("/", handler.TestcaseModify)
		testcaseEditorRoute.DELETE("/:id", handler.TestcaseRemove)
		testcaseEditorRoute.POST("/datamake", handler.TestcaseDataMake)
	}
}

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
