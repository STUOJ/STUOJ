package bootstrap

import (
	"STUOJ/internal/infrastructure/neko"
	"STUOJ/pkg/config"
	"log"
)

func initNeko() {
	err := neko.InitNekoAcm(config.Conf.NekoAcm.Host, config.Conf.NekoAcm.Port, config.Conf.NekoAcm.Token)
	if err != nil {
		log.Println(err)
		log.Println("初始化 NekoACM 失败！")
		return
	}
	log.Println("初始化 NekoACM 成功")
}
