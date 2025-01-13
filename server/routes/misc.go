package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitMiscRoute(ginServer *gin.Engine) {
	uploadUserRoute := ginServer.Group("/upload")
	{
		// 使用中间件
		uploadUserRoute.Use(middlewares.TokenAuthUser())

		uploadUserRoute.POST("/image", handler.UploadImage)
	}

	emailPublicRoute := ginServer.Group("/email")
	{
		emailPublicRoute.POST("/send", handler.SendVerificationCode)
	}
}
