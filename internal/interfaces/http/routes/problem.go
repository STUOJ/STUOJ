package routes

import (
	handler2 "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitProblemRoute(ginServer *gin.Engine) {
	problemPublicRoute := ginServer.Group("/problem")
	{
		problemPublicRoute.GET("/", handler2.ProblemList)
		problemPublicRoute.GET("/:id", handler2.ProblemInfo)
	}

	problemEditorRoute := ginServer.Group("/problem")
	{
		// 使用中间件
		problemEditorRoute.Use(middlewares.TokenAuthEditor())

		problemEditorRoute.POST("/", handler2.ProblemAdd)
		problemEditorRoute.PUT("/", handler2.ProblemModify)

		problemEditorRoute.GET("/history/", handler2.HistoryListOfProblem)
		problemEditorRoute.GET("/history/:id", handler2.HistoryInfo)
	}
}

func InitTagRoute(ginServer *gin.Engine) {
	tagPublicRoute := ginServer.Group("/tag")
	{
		tagPublicRoute.GET("/", handler2.TagList)
		tagPublicRoute.GET("/:id", handler2.TagInfo)
	}

	tagEditorRoute := ginServer.Group("/tag")
	{
		// 使用中间件
		tagEditorRoute.Use(middlewares.TokenAuthEditor())

		tagEditorRoute.POST("/", handler2.TagAdd)
		tagEditorRoute.PUT("/", handler2.TagModify)
		tagEditorRoute.DELETE("/:id", handler2.TagRemove)
	}
}

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthEditor())

		testcaseEditorRoute.GET("/", handler2.TestcaseList)
		testcaseEditorRoute.GET("/:id", handler2.TestcaseInfo)
		testcaseEditorRoute.POST("/", handler2.TestcaseAdd)
		testcaseEditorRoute.PUT("/", handler2.TestcaseModify)
		testcaseEditorRoute.DELETE("/:id", handler2.TestcaseRemove)
		testcaseEditorRoute.POST("/datamake", handler2.TestcaseDataMake)
	}
}

func InitSolutionRoute(ginServer *gin.Engine) {
	solutionEditorRoute := ginServer.Group("/solution")
	{
		// 使用中间件
		solutionEditorRoute.Use(middlewares.TokenAuthEditor())

		solutionEditorRoute.GET("/", handler2.SolutionList)
		solutionEditorRoute.GET("/:id", handler2.SolutionInfo)
		solutionEditorRoute.POST("/", handler2.SolutionAdd)
		solutionEditorRoute.PUT("/", handler2.SolutionModify)
		solutionEditorRoute.DELETE("/:id", handler2.SolutionRemove)
	}
}

func InitCollectionRoute(ginServer *gin.Engine) {
	collectionPublicRoute := ginServer.Group("/collection")
	{
		collectionPublicRoute.GET("/:id", handler2.CollectionInfo)
		collectionPublicRoute.GET("/", handler2.CollectionList)
	}

	collectionUserRoute := ginServer.Group("/collection")
	{
		// 使用中间件
		collectionUserRoute.Use(middlewares.TokenAuthUser())

		collectionUserRoute.POST("/", handler2.CollectionAdd)
		collectionUserRoute.PUT("/", handler2.CollectionModify)
		collectionUserRoute.DELETE("/:id", handler2.CollectionRemove)

		collectionUserRoute.PUT("/problem", handler2.CollectionModifyProblem)

		collectionUserRoute.PUT("/user", handler2.CollectionModifyUser)
	}
}
