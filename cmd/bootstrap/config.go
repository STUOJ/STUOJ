package bootstrap

import (
	"STUOJ/pkg/config"
	"log"
)

func InitConfig() {
	err := config.InitConfig()
	if err != nil {
		log.Println("初始化配置失败！")
		panic(err)
	}
	log.Println("初始化配置成功")
}
