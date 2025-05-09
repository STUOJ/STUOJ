package routes

import (
	handler "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitProblemRoute(ginServer *gin.Engine) {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", handler.ProblemList)
		problemPublicRoute.GET("/:id", handler.ProblemInfo)
		problemPublicRoute.GET("/statistics", handler.ProblemStatistics)
	}

	problemEditorRoute := ginServer.Group("/problem")
	{
		// 使用中间件
		problemEditorRoute.Use(middlewares.TokenAuthEditor())

		problemEditorRoute.POST("/", handler.ProblemAdd)
		problemEditorRoute.PUT("/", handler.ProblemModify)

		problemEditorRoute.GET("/history/", handler.HistoryListOfProblem)
		problemEditorRoute.GET("/history/:id", handler.HistoryInfo)
	}
}

func InitTagRoute(ginServer *gin.Engine) {
	tagPublicRoute := ginServer.Group("/tag")
	{
		tagPublicRoute.GET("/", handler.TagList)
		tagPublicRoute.GET("/:id", handler.TagInfo)
	}

	tagEditorRoute := ginServer.Group("/tag")
	{
		// 使用中间件
		tagEditorRoute.Use(middlewares.TokenAuthEditor())

		tagEditorRoute.POST("/", handler.TagAdd)
		tagEditorRoute.PUT("/", handler.TagModify)
		tagEditorRoute.DELETE("/:id", handler.TagRemove)
	}
}

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthEditor())

		testcaseEditorRoute.GET("/", handler.TestcaseList)
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
		solutionEditorRoute.Use(middlewares.TokenAuthEditor())

		solutionEditorRoute.GET("/", handler.SolutionList)
		solutionEditorRoute.GET("/:id", handler.SolutionInfo)
		solutionEditorRoute.POST("/", handler.SolutionAdd)
		solutionEditorRoute.PUT("/", handler.SolutionModify)
		solutionEditorRoute.DELETE("/:id", handler.SolutionRemove)
		solutionEditorRoute.GET("/statistics", handler.SolutionStatistics)
	}
}

func InitCollectionRoute(ginServer *gin.Engine) {
	collectionPublicRoute := ginServer.Group("/collection")
	{
		collectionPublicRoute.GET("/:id", handler.CollectionInfo)
		collectionPublicRoute.GET("/", handler.CollectionList)
	}

	collectionUserRoute := ginServer.Group("/collection")
	{
		// 使用中间件
		collectionUserRoute.Use(middlewares.TokenAuthUser())

		collectionUserRoute.POST("/", handler.CollectionAdd)
		collectionUserRoute.PUT("/", handler.CollectionModify)
		collectionUserRoute.DELETE("/:id", handler.CollectionRemove)

		collectionUserRoute.PUT("/problem", handler.CollectionModifyProblem)

		collectionUserRoute.PUT("/user", handler.CollectionModifyUser)
	}
}
