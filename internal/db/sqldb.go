package db

import (
	"STUOJ/internal/conf"
	"time"
)

func createSqlDb() error {
	var err error
	config := conf.Conf.Datebase

	SqlDb, err = Db.DB()
	if err != nil {
		return err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	SqlDb.SetMaxIdleConns(config.MaxIdle)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	SqlDb.SetMaxOpenConns(config.MaxConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	SqlDb.SetConnMaxLifetime(time.Hour)

	err = SqlDb.Ping()
	if err != nil {
		return err
	}

	return nil
}