package http

import (
	middlewares2 "STUOJ/internal/interfaces/http/middlewares"
	routes2 "STUOJ/internal/interfaces/http/routes"
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
	ginServer.Use(middlewares2.TokenGetInfo())
	ginServer.Use(middlewares2.ErrorHandler())
	ginServer.Use(middlewares2.QueryCleaner())

	// 初始化路由

	// routes/user.go
	routes2.InitUserRoute(ginServer)

	// routes/problem.go
	routes2.InitProblemRoute(ginServer)
	routes2.InitTagRoute(ginServer)
	routes2.InitSolutionRoute(ginServer)
	routes2.InitTestcaseRoute(ginServer)
	routes2.InitCollectionRoute(ginServer)

	// routes/judge.go
	routes2.InitJudgeRoute(ginServer)
	routes2.InitRecordRoute(ginServer)
	routes2.InitLanguageRoute(ginServer)

	// routes/contest.go
	//routes2.InitContestRoute(ginServer)
	//routes2.InitTeamRoute(ginServer)

	// routes/blog.go
	routes2.InitBlogRoute(ginServer)
	routes2.InitCommentRoute(ginServer)

	// routes/system.go
	routes2.InitSystemRoute(ginServer)

	// routes/misc.go
	routes2.InitAiRouter(ginServer)
	routes2.InitMiscRoute(ginServer)

	return nil
}
