package bootstrap

import (
	"STUOJ/external/yuki"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println(err)
		log.Println("初始化 yuki-image 失败！")
		return
	}
	log.Println("初始化 yuki-image 成功")
}
