package solution

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除题解
func Delete(id uint64) error {
	// 查询题解
	_, err := dao.SelectSolutionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题解不存在")
	}

	// 删除题解
	err = dao.DeleteSolutionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除题解失败")
	}

	return nil
}
