package http

import (
	middlewares "STUOJ/internal/interfaces/http/middlewares"
	routes "STUOJ/internal/interfaces/http/routes"
	"STUOJ/internal/interfaces/http/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, vo.RespOk("STUOJ后端启动成功！", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, vo.RespError("404 Not Found", nil))
	})

	// 使用中间件
	ginServer.Use(middlewares.TokenGetInfo())
	ginServer.Use(middlewares.ErrorHandler())
	ginServer.Use(middlewares.QueryCleaner())

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
	//routes.InitContestRoute(ginServer)
	//routes.InitTeamRoute(ginServer)

	// routes/blog.go
	routes.InitBlogRoute(ginServer)
	routes.InitCommentRoute(ginServer)

	// routes/system.go
	routes.InitSystemRoute(ginServer)

	// routes/misc.go
	routes.InitAiRouter(ginServer)
	routes.InitMiscRoute(ginServer)

	return nil
}
