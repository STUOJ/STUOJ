package tag

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除标签
func Delete(id uint64) error {
	// 查询标签
	_, err := dao.SelectTagById(id)
	if err != nil {
		log.Println(err)
		return errors.New("标签不存在")
	}

	// 删除标签
	err = dao.DeleteTagById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除标签失败")
	}

	return nil
}
