package routes

import (
	handler2 "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBlogRoute(ginServer *gin.Engine) {
	blogPublicRoute := ginServer.Group("/blog")
	{
		blogPublicRoute.GET("/", handler2.BlogList)
		blogPublicRoute.GET("/:id", handler2.BlogInfo)
	}

	blogUserRoute := ginServer.Group("/blog")
	{
		// 使用中间件
		blogUserRoute.Use(middlewares.TokenAuthUser())

		blogUserRoute.POST("/", handler2.BlogUpload)
		blogUserRoute.PUT("/", handler2.BlogEdit)
		blogUserRoute.DELETE("/:id", handler2.BlogRemove)
	}
}

func InitCommentRoute(ginServer *gin.Engine) {
	commentPublicRoute := ginServer.Group("/comment")
	{
		commentPublicRoute.GET("/", handler2.CommentList)
	}

	commentUserRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentUserRoute.Use(middlewares.TokenAuthUser())

		commentUserRoute.POST("/", handler2.CommentAdd)
	}

	commentAdminRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentAdminRoute.Use(middlewares.TokenAuthAdmin())

		commentAdminRoute.PUT("/", handler2.CommentModify)
	}
}
