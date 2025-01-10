package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAiRouter(ginServer *gin.Engine) {
	aiUserRouter := ginServer.Group("/ai")
	{
		aiUserRouter.Use(middlewares.TokenAuthUser())

		aiUserRouter.POST("/chat/assistant", handler.ChatAssistant)
		aiUserRouter.GET("/misc/joke", handler.TellJoke)
		aiUserRouter.POST("/judge/submit", handler.SubmitVirtualJudge)
	}

	aiEditorRouter := ginServer.Group("/ai")
	{
		aiEditorRouter.Use(middlewares.TokenAuthEditor())

		aiEditorRouter.POST("/problem/parse", handler.ParseProblem)
		aiEditorRouter.POST("/problem/translate", handler.TranslateProblem)
		aiEditorRouter.POST("/problem/generate", handler.GenerateProblem)
		aiEditorRouter.POST("/testcase/generate", handler.GenerateTestcase)
		aiEditorRouter.POST("/solution/generate", handler.GenerateSolution)
	}
}
