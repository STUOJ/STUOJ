package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitMiscRoute(ginServer *gin.Engine) {
	uploadRoute := ginServer.Group("/upload")
	{
		uploadRoute.Use(middlewares.TokenAuthUser())
		uploadRoute.POST("/image", handler.UploadImage)
	}

	miscRoute := ginServer.Group("/misc")
	{
		miscRoute.GET("/joke", handler.GetJoke)
	}
}
