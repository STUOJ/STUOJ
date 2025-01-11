package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthAdmin())

		testcaseEditorRoute.GET("/:id", handler.TestcaseInfo)
		testcaseEditorRoute.POST("/", handler.TestcaseAdd)
		testcaseEditorRoute.PUT("/", handler.TestcaseModify)
		testcaseEditorRoute.DELETE("/:id", handler.TestcaseRemove)
		testcaseEditorRoute.POST("/datamake", handler.TestcaseDataMake)
	}
}
