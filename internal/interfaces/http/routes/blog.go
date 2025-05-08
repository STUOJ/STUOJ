package routes

import (
	handler "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBlogRoute(ginServer *gin.Engine) {
	blogPublicRoute := ginServer.Group("/blog")
	{
		blogPublicRoute.GET("/", handler.BlogList)
		blogPublicRoute.GET("/:id", handler.BlogInfo)
		blogPublicRoute.GET("/statistics", handler.BlogStatistics)
	}

	blogUserRoute := ginServer.Group("/blog")
	{
		// 使用中间件
		blogUserRoute.Use(middlewares.TokenAuthUser())

		blogUserRoute.POST("/", handler.BlogUpload)
		blogUserRoute.PUT("/", handler.BlogEdit)
		blogUserRoute.DELETE("/:id", handler.BlogRemove)
	}
}

func InitCommentRoute(ginServer *gin.Engine) {
	commentPublicRoute := ginServer.Group("/comment")
	{
		commentPublicRoute.GET("/", handler.CommentList)
	}

	commentUserRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentUserRoute.Use(middlewares.TokenAuthUser())

		commentUserRoute.POST("/", handler.CommentAdd)
		commentUserRoute.GET("/statistics", handler.CommentStatistics)
	}

	commentAdminRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentAdminRoute.Use(middlewares.TokenAuthAdmin())

		commentAdminRoute.PUT("/", handler.CommentModify)
	}
}
