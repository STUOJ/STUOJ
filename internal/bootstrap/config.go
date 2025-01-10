package bootstrap

import (
	"STUOJ/internal/conf"
	"log"
)

func initConfig() {
	err := conf.InitConfig()
	if err != nil {
		log.Println("初始化配置失败！")
		panic(err)
	}
	log.Println("初始化配置成功")
}
