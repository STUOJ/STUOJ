package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

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
		commentUserRoute.DELETE("/:id", handler.CommentRemove)
	}
	commentAdminRoute := ginServer.Group("/comment")
	{
		// 使用中间件
		commentAdminRoute.Use(middlewares.TokenAuthAdmin())

		commentAdminRoute.PUT("/", handler.CommentModify)
	}
}
