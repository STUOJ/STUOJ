package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入题单
func Insert(c entity.Collection) (uint64, error) {
	var err error

	updateTime := time.Now()
	c.UpdateTime = updateTime
	c.CreateTime = updateTime

	// 插入题单
	c.Id, err = dao.InsertCollection(c)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入题单失败")
	}

	return c.Id, nil
}
