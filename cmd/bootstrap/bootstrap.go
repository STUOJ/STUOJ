package bootstrap

func Init() {
	initConfig()
	initDatabase()

	// 异步初始化
	go initJudge0()
	go initYuki()
	go initNeko()

	initServer()
}
