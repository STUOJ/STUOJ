package routes

import (
	"STUOJ/server/handler"
	"github.com/gin-gonic/gin"
)

func InitTagRoute(ginServer *gin.Engine) {
	tagPublicRoute := ginServer.Group("/tag")
	{
		tagPublicRoute.GET("/", handler.TagList)
	}
	tagEditorRoute := ginServer.Group("/tag")
	{
		tagEditorRoute.POST("/", handler.TagAdd)
		tagEditorRoute.PUT("/", handler.TagModify)
		tagEditorRoute.DELETE("/:id", handler.TagRemove)
	}
}
