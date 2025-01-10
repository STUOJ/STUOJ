package bootstrap

import (
	"STUOJ/internal/db"
	"log"
)

func initDatabase() {
	err := db.InitDatabase()
	if err != nil {
		log.Println("初始化数据库失败！")
		panic(err)
	}
	log.Println("初始化数据库成功")
}
