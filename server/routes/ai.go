package routes

import (
	"STUOJ/external/neko"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAiRouter(ginServer *gin.Engine) {
	aiUserRouter := ginServer.Group("/ai")
	{
		// 使用中间件
		aiUserRouter.Use(middlewares.TokenAuthUser())

		aiUserRouter.POST("/chat/assistant", neko.ForwardHandler)
		aiUserRouter.GET("/misc/joke", neko.ForwardHandler)
		aiUserRouter.POST("/judge/submit", neko.ForwardHandler)
	}

	aiEditorRouter := ginServer.Group("/ai")
	{
		// 使用中间件
		aiEditorRouter.Use(middlewares.TokenAuthEditor())

		aiEditorRouter.POST("/problem/parse", neko.ForwardHandler)
		aiEditorRouter.POST("/problem/translate", neko.ForwardHandler)
		aiEditorRouter.POST("/problem/generate", neko.ForwardHandler)
		aiEditorRouter.POST("/testcase/generate", neko.ForwardHandler)
		aiEditorRouter.POST("/solution/generate", neko.ForwardHandler)
	}
}
