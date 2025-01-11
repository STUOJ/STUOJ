package routes

import (
	"STUOJ/server/handler-admin"
	"STUOJ/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAdminRoute(ginServer *gin.Engine) {
	adminPrivateRoute := ginServer.Group("/admin")
	{
		// 使用中间件
		adminPrivateRoute.Use(middlewares.TokenAuthAdmin())

		{
			adminPrivateRoute.GET("/user", handler_admin.UserList)
			adminPrivateRoute.POST("/user", handler_admin.AdminUserAdd)
			adminPrivateRoute.DELETE("/user/:id", handler_admin.AdminUserRemove)
		}
		{
			adminPrivateRoute.GET("/testcase/:id", handler_admin.AdminTestcaseInfo)
			adminPrivateRoute.POST("/testcase", handler_admin.AdminTestcaseAdd)
			adminPrivateRoute.PUT("/testcase", handler_admin.AdminTestcaseModify)
			adminPrivateRoute.DELETE("/testcase/:id", handler_admin.AdminTestcaseRemove)
			adminPrivateRoute.POST("/testcase/datamake", handler_admin.AdminTestcaseDataMake)
		}
		{
			adminPrivateRoute.GET("/tag", handler_admin.AdminTagList)
			adminPrivateRoute.POST("/tag", handler_admin.AdminTagAdd)
			adminPrivateRoute.PUT("/tag", handler_admin.AdminTagModify)
			adminPrivateRoute.DELETE("/tag/:id", handler_admin.AdminTagRemove)
		}
	}
}
