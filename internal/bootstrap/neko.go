package bootstrap

import (
	"STUOJ/external/neko"
	"STUOJ/internal/conf"
	"log"
)

func initNeko() {
	err := neko.InitNekoAcm(conf.Conf.NekoAcm.Host, conf.Conf.NekoAcm.Port, conf.Conf.NekoAcm.Token)
	if err != nil {
		log.Println(err)
		log.Println("初始化 NekoACM 失败！")
		return
	}
	log.Println("初始化 NekoACM 成功")
}
