package routes

import (
	"STUOJ/server/handler"
	"github.com/gin-gonic/gin"
)

func InitRecordRoute(ginServer *gin.Engine) {
	recordPublicRoute := ginServer.Group("/record")
	{
		recordPublicRoute.GET("/", handler.RecordList)
		recordPublicRoute.GET("/:id", handler.RecordInfo)
		recordPublicRoute.GET("/user/:id", handler.RecordListOfUser)
		recordPublicRoute.GET("/problem/:id", handler.RecordListOfProblem)
	}
}