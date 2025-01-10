package bootstrap

import (
	"STUOJ/server"
	"log"
)

func initServer() {
	err := server.InitServer()
	if err != nil {
		log.Println("初始化服务器失败！")
		panic(err)
	}
	log.Println("初始化服务器成功")
}
