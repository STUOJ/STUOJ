package db

import (
	"STUOJ/internal/conf"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库
func InitDatabase() error {
	var err error
	config := conf.Conf.Datebase

	gormConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	}

	dsn := config.User + ":" + config.Pwd + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println("正在连接数据库：", dsn)
	Db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Println("连接数据库失败！")
		return err
	}

	err = createSqlDb()
	if err != nil {
		log.Println("连接数据库失败！")
		return err
	}

	// autoMigrate()

	return nil
}
