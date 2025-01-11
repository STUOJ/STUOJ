package routes

import (
	"STUOJ/server/handler"
	"github.com/gin-gonic/gin"
)

func InitStatisticsRoute(ginServer *gin.Engine) {
	statisticsPublicRoute := ginServer.Group("/statistics")
	{
		statisticsPublicRoute.GET("/user", handler.StatisticsUser)
		statisticsPublicRoute.GET("/user/role", handler.StatisticsUserOfRole)
		statisticsPublicRoute.GET("/user/register", handler.StatisticsUserOfRegister)

		statisticsPublicRoute.GET("/tag", handler.StatisticsTag)
		statisticsPublicRoute.GET("/problem", handler.StatisticsProblem)
		statisticsPublicRoute.GET("/problem/insert", handler.StatisticsProblemOfInsert)
		statisticsPublicRoute.GET("/problem/update", handler.StatisticsProblemOfUpdate)
		statisticsPublicRoute.GET("/problem/delete", handler.StatisticsProblemOfDelete)

		statisticsPublicRoute.GET("/judge", handler.StatisticsJudge)

		statisticsPublicRoute.GET("/record", handler.StatisticsRecord)
		statisticsPublicRoute.GET("/record/submit", handler.StatisticsRecordOfSubmit)
		statisticsPublicRoute.GET("/record/language", handler.StatisticsRecordOfLanguage)
		statisticsPublicRoute.GET("/submission/status", handler.StatisticsSubmissionOfStatus)
		statisticsPublicRoute.GET("/judgement/status", handler.StatisticsJudgementOfStatus)

		statisticsPublicRoute.GET("/blog", handler.StatisticsBlog)
		statisticsPublicRoute.GET("/blog/submit", handler.StatisticsBlogOfSubmit)
		statisticsPublicRoute.GET("/comment/submit", handler.StatisticsCommentOfSubmit)
	}
}
