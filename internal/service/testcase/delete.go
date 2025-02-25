package testcase

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除评测点数据
func Delete(id uint64) error {
	// 查询评测点
	_, err := dao.SelectTestcaseById(id)
	if err != nil {
		log.Println(err)
		return errors.New("评测点不存在")
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(id)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 删除评测点
	err = dao.DeleteTestcaseById(id)
	if err != nil {
		return err
	}

	return nil
}
