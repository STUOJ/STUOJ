package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 插入题单
func Insert(c entity.Collection) (uint64, error) {
	var err error

	c.Id, err = dao.InsertCollection(c)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败")
	}

	return c.Id, nil
}
