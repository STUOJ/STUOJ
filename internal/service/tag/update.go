package tag

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID更新标签
func Update(id uint64, n string) error {
	// 查询标签
	t, err := dao.SelectTagById(id)
	if err != nil {
		log.Println(err)
		return errors.New("标签不存在")
	}

	// 更新标签名
	t.Name = n

	// 更新标签
	err = dao.UpdateTagById(t)
	if err != nil {
		log.Println(err)
		return errors.New("修改失败，标签名不能重复")
	}

	return nil
}
