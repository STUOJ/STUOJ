package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBlogRoute(ginServer *gin.Engine) {
	blogPublicRoute := ginServer.Group("/blog")
	{
		blogPublicRoute.GET("/", handler.BlogList)
		blogPublicRoute.GET("/:id", handler.BlogInfo)
	}

	blogUserRoute := ginServer.Group("/blog")
	{
		// 使用中间件
		blogUserRoute.Use(middlewares.TokenAuthUser())

		blogUserRoute.POST("/", handler.BlogUpload)
		blogUserRoute.PUT("/", handler.BlogEdit)
		blogUserRoute.PUT("/:id", handler.BlogSubmit)
		blogUserRoute.DELETE("/:id", handler.BlogRemove)
	}
}
