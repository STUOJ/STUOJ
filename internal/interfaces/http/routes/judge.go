package routes

import (
	handler "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitJudgeRoute(ginServer *gin.Engine) {
	judgeUserRoute := ginServer.Group("/judge")
	{
		// 使用中间件
		judgeUserRoute.Use(middlewares.TokenAuthUser())

		judgeUserRoute.POST("/submit", handler.JudgeSubmit)
		judgeUserRoute.POST("/testrun", handler.JudgeTestRun)
	}
}

func InitRecordRoute(ginServer *gin.Engine) {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", handler.RecordList)
		recordPublicRoute.GET("/:id", handler.RecordInfo)
		recordPublicRoute.GET("/ac/user", handler.SelectACUsers)
		recordPublicRoute.GET("/statistics", handler.RecordStatistics)
	}

	recordAdminRoute := ginServer.Group("/record")
	{
		// 使用中间件
		recordAdminRoute.Use(middlewares.TokenAuthAdmin())
	}
}

func InitLanguageRoute(ginServer *gin.Engine) {
	languagePublicRouter := ginServer.Group("/language")
	{
		languagePublicRouter.GET("/list", handler.LanguageList)
	}

	languageAdminRouter := ginServer.Group("/language")
	{
		// 使用中间件
		languageAdminRouter.Use(middlewares.TokenAuthAdmin())

		languageAdminRouter.PUT("/update", handler.UpdateLanguage)
		languageAdminRouter.GET("/statistics", handler.LanguageStatistics)
	}
}
