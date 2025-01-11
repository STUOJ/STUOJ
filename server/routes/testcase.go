package routes

import (
	"STUOJ/server/handler"
	"STUOJ/server/handler-admin"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitTestcaseRoute(ginServer *gin.Engine) {
	testcaseEditorRoute := ginServer.Group("/testcase")
	{
		// 使用中间件
		testcaseEditorRoute.Use(middlewares.TokenAuthAdmin())

		{
			testcaseEditorRoute.GET("/user", handler_admin.UserList)
			testcaseEditorRoute.POST("/user", handler_admin.AdminUserAdd)
			testcaseEditorRoute.DELETE("/user/:id", handler_admin.AdminUserRemove)
		}
		{
			testcaseEditorRoute.GET("/:id", handler.TestcaseInfo)
			testcaseEditorRoute.POST("/", handler.TestcaseAdd)
			testcaseEditorRoute.PUT("/", handler.TestcaseModify)
			testcaseEditorRoute.DELETE("/:id", handler.TestcaseRemove)
			testcaseEditorRoute.POST("/datamake", handler.TestcaseDataMake)
		}
	}
}
