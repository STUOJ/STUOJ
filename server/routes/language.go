package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitLanguageRoute(ginServer *gin.Engine) {
	languagePublicRouter := ginServer.Group("/language")
	{
		languagePublicRouter.GET("/list", handler.LanguageList)
	}

	languageAdminRouter := ginServer.Group("/language")
	{
		// 使用中间件
		languageAdminRouter.Use(middlewares.TokenAuthAdmin())

		languageAdminRouter.PUT("/update", handler.UpdateLanguage)
	}
}
