package server

import (
	"STUOJ/internal/conf"
	"STUOJ/server/validator"

	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func InitServer() error {
	config := conf.Conf.Server

	// 创建gin实例
	ginServer = gin.Default()

	validator.SetupValidator()

	// 初始化路由
	err := InitRoute()
	if err != nil {
		return err
	}

	// 启动服务
	err = ginServer.Run(":" + config.Port)
	if err != nil {
		return err
	}

	return nil
}
