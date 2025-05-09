package interfaces_test

import (
	"STUOJ/internal/infrastructure/persistence/repository"
	"STUOJ/pkg/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	var err error

	err = os.Chdir("../../..")
	if err != nil {
		panic("切换工作目录失败：" + err.Error())
	}

	err = config.InitConfig()
	if err != nil {
		panic("配置文件加载失败：" + err.Error())
	}

	err = repository.InitDatabase()
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}

	code := m.Run()

	os.Exit(code)
}
