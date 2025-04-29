package bootstrap

import (
	"STUOJ/internal/interfaces/http"
	"log"
)

func initServer() {
	err := http.InitServer()
	if err != nil {
		log.Println("初始化服务器失败！")
		panic(err)
	}
	log.Println("初始化服务器成功")
}
