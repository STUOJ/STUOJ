package bootstrap

import (
	"STUOJ/internal/infrastructure/persistence/repository"
	"log"
)

func InitDatabase() {
	err := repository.InitDatabase()
	if err != nil {
		log.Println("初始化数据库失败！")
		panic(err)
	}
	log.Println("初始化数据库成功")
}
