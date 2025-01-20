package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除题单
func Delete(id uint64, uid uint64, role entity.Role) error {
	// 查询题单
	c0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("题单不存在")
	}

	if role < entity.RoleAdmin {
		if c0.UserId != uid {
			return errors.New("没有权限，只能删除自己的题单")
		}
	}

	// 删除题单
	err = dao.DeleteCollectionById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除题单失败")
	}

	return nil
}
