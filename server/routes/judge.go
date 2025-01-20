package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitJudgeRoute(ginServer *gin.Engine) {
	judgePublicRoute := ginServer.Group("/judge")
	{
		judgePublicRoute.GET("/language", handler.JudgeLanguageList)
	}

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
	}

	recordAdminRoute := ginServer.Group("/record")
	{
		// 使用中间件
		recordAdminRoute.Use(middlewares.TokenAuthAdmin())

		recordAdminRoute.DELETE("/:id", handler.RecordRemove)
	}
}
