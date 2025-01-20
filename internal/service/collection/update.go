package collection

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID更新题单
func UpdateById(id uint64, n string) error {
	// 查询题单
	t, err := dao.SelectCollectionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	// 更新题单

	// 更新题单
	err = dao.UpdateCollectionById(t)
	if err != nil {
		log.Println(err)
		return errors.New("修改失败")
	}

	return nil
}
