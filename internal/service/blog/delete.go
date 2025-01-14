package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除博客（检查用户ID）
func Delete(id uint64, uid uint64, role entity.Role) error {
	// 查询博客
	b0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("博客不存在")
	}

	if role < entity.RoleAdmin {
		if b0.UserId != uid {
			return errors.New("没有权限，只能删除自己的博客")
		}
		if b0.Status == entity.BlogBanned {
			return errors.New("该博客已被封禁，无法删除")
		}
		if b0.Status == entity.BlogNotice {
			return errors.New("只有管理员才能删除公告")
		}
	}

	// 删除博客
	err = dao.DeleteBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除博客失败")
	}

	return nil
}
