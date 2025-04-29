package bootstrap

import (
	"STUOJ/internal/infrastructure/repository"
	"log"
)

func initDatabase() {
	err := repository.InitDatabase()
	if err != nil {
		log.Println("初始化数据库失败！")
		panic(err)
	}
	log.Println("初始化数据库成功")
}
