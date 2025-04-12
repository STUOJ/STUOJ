package server

import (
	"STUOJ/internal/model"
	"STUOJ/server/middlewares"
	"STUOJ/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("STUOJ后端启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 使用中间件
	ginServer.Use(middlewares.TokenGetInfo())
	ginServer.Use(middlewares.ErrorHandler())

	// 初始化路由

	// routes/user.go
	routes.InitUserRoute(ginServer)

	// routes/problem.go
	routes.InitProblemRoute(ginServer)
	routes.InitTagRoute(ginServer)
	routes.InitSolutionRoute(ginServer)
	routes.InitTestcaseRoute(ginServer)
	routes.InitCollectionRoute(ginServer)

	// routes/judge.go
	routes.InitJudgeRoute(ginServer)
	routes.InitRecordRoute(ginServer)
	routes.InitLanguageRoute(ginServer)

	// routes/contest.go
	routes.InitContestRoute(ginServer)
	routes.InitTeamRoute(ginServer)

	// routes/blog.go
	routes.InitBlogRoute(ginServer)
	routes.InitCommentRoute(ginServer)

	// routes/statistics.go
	routes.InitStatisticsRoute(ginServer)

	// routes/system.go
	routes.InitSystemRoute(ginServer)

	// routes/misc.go
	routes.InitAiRouter(ginServer)
	routes.InitMiscRoute(ginServer)

	return nil
}
