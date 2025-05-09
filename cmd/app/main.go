package app

import "STUOJ/cmd/bootstrap"

func Main() {
	bootstrap.InitConfig()
	bootstrap.InitDatabase()

	// 异步初始化
	go bootstrap.InitJudge0()
	go bootstrap.InitYuki()
	go bootstrap.InitNeko()

	bootstrap.InitServer()
}
