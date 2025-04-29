package routes

import (
	"STUOJ/http/handler"
	handler2 "STUOJ/internal/interfaces/http/handler"
	"STUOJ/internal/interfaces/http/middlewares"
	"github.com/gin-gonic/gin"
)

func InitContestRoute(ginServer *gin.Engine) {
	contestPublicRoute := ginServer.Group("/contest")
	{
		contestPublicRoute.GET("/:id", handler2.ContestInfo)
		contestPublicRoute.GET("/", handler2.ContestList)
	}

	contestAdminRoute := ginServer.Group("/contest")
	{
		// 使用中间件
		contestAdminRoute.Use(middlewares.TokenAuthAdmin())

		contestAdminRoute.POST("/", handler2.ContestAdd)
		contestAdminRoute.PUT("/", handler2.ContestModify)
		contestAdminRoute.DELETE("/:id", handler2.ContestRemove)
	}
}

func InitTeamRoute(ginServer *gin.Engine) {
	teamPublicRoute := ginServer.Group("/team")
	{
		teamPublicRoute.GET("/:id", handler.TeamInfo)
		teamPublicRoute.GET("/", handler.TeamList)
	}

	teamUserRoute := ginServer.Group("/team")
	{
		// 使用中间件
		teamUserRoute.Use(middlewares.TokenAuthUser())

		teamUserRoute.POST("/", handler.TeamAdd)
		teamUserRoute.PUT("/", handler.TeamModify)
		teamUserRoute.DELETE("/:id", handler.TeamRemove)
	}
}
