package collection

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除题单
func DeleteById(id uint64) error {
	// 查询题单
	_, err := dao.SelectCollectionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	// 删除题单
	err = dao.DeleteCollectionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除失败")
	}

	return nil
}
