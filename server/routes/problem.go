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
		problemEditorRoute.DELETE("/:id", handler.ProblemRemove)

		problemEditorRoute.POST("/tag", handler.ProblemAddTag)
		problemEditorRoute.DELETE("/tag", handler.ProblemRemoveTag)
		problemEditorRoute.GET("/history/:id", handler.HistoryListOfProblem)
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
		tagEditorRoute.DELETE("/:id", handler.TagRemove)
	}
}

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthEditor())

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

		solutionEditorRoute.GET("/:id", handler.SolutionInfo)
		solutionEditorRoute.POST("/", handler.SolutionAdd)
		solutionEditorRoute.PUT("/", handler.SolutionModify)
		solutionEditorRoute.DELETE("/:id", handler.SolutionRemove)
	}
}

func InitCollectionRoute(ginServer *gin.Engine) {
	collectionPublicRoute := ginServer.Group("/collection")
	{
		collectionPublicRoute.GET("/", handler.CollectionList)
		collectionPublicRoute.GET("/:id", handler.CollectionInfo)
	}

	collectionUserRoute := ginServer.Group("/collection")
	{
		// 使用中间件
		collectionUserRoute.Use(middlewares.TokenAuthUser())

		collectionUserRoute.POST("/", handler.CollectionAdd)
		collectionUserRoute.PUT("/", handler.CollectionModify)
		collectionUserRoute.DELETE("/:id", handler.CollectionRemove)

		collectionUserRoute.POST("/problem", handler.CollectionAddProblem)
		collectionUserRoute.PUT("/problem", handler.CollectionModifyProblem)
		collectionUserRoute.DELETE("/problem", handler.CollectionRemoveProblem)
	}
}
