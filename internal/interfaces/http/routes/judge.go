package routes

import (
	handler2 "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitJudgeRoute(ginServer *gin.Engine) {
	judgeUserRoute := ginServer.Group("/judge")
	{
		// 使用中间件
		judgeUserRoute.Use(middlewares.TokenAuthUser())

		judgeUserRoute.POST("/submit", handler2.JudgeSubmit)
		judgeUserRoute.POST("/testrun", handler2.JudgeTestRun)
	}
}

func InitRecordRoute(ginServer *gin.Engine) {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", handler2.RecordList)
		recordPublicRoute.GET("/:id", handler2.RecordInfo)
		recordPublicRoute.GET("/ac/user", handler2.SelectACUsers)
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
		languagePublicRouter.GET("/list", handler2.LanguageList)
	}

	languageAdminRouter := ginServer.Group("/language")
	{
		// 使用中间件
		languageAdminRouter.Use(middlewares.TokenAuthAdmin())

		languageAdminRouter.PUT("/update", handler2.UpdateLanguage)
	}
}
