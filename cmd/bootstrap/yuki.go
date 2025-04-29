package bootstrap

import (
	"STUOJ/internal/infrastructure/yuki"
	"STUOJ/pkg/config"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage(config.Conf.YukiImage.Host, config.Conf.YukiImage.Port, config.Conf.YukiImage.Token)
	if err != nil {
		log.Println(err)
		log.Println("初始化 yuki-image 失败！")
		return
	}
	log.Println("初始化 yuki-image 成功")
}
