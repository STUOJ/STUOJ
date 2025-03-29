package bootstrap

import (
	"STUOJ/external/yuki"
	"STUOJ/internal/conf"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage(conf.Conf.YukiImage.Host, conf.Conf.YukiImage.Port, conf.Conf.YukiImage.Token)
	if err != nil {
		log.Println(err)
		log.Println("初始化 yuki-image 失败！")
		return
	}
	log.Println("初始化 yuki-image 成功")
}
