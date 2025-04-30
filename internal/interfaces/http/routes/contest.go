package routes

/*
func InitContestRoute(ginServer *gin.Engine) {
	contestPublicRoute := ginServer.Group("/contest")
	{
		contestPublicRoute.GET("/:id", handler.ContestInfo)
		contestPublicRoute.GET("/", handler.ContestList)
	}

	contestAdminRoute := ginServer.Group("/contest")
	{
		// 使用中间件
		contestAdminRoute.Use(middlewares.TokenAuthAdmin())

		contestAdminRoute.POST("/", handler.ContestAdd)
		contestAdminRoute.PUT("/", handler.ContestModify)
		contestAdminRoute.DELETE("/:id", handler.ContestRemove)
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
*/
