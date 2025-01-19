package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除评论
func DeleteById(id uint64) error {
	// 查询评论
	_, err := dao.SelectCommentById(id)
	if err != nil {
		log.Println(err)
		return errors.New("评论不存在")
	}

	// 删除评论
	err = dao.DeleteCommentById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除评论失败")
	}

	return nil
}

// 根据ID删除评论（检查用户ID）
func DeleteByIdCheckUserId(id uint64, uid uint64, role entity.Role) error {
	// 查询评论
	c0, err := dao.SelectCommentById(id)
	if err != nil {
		log.Println(err)
		return errors.New("评论不存在")
	}

	// 检查权限
	log.Println(c0.UserId, uid)
	if c0.UserId != uid && role < entity.RoleAdmin {
		return errors.New("没有权限，只能删除自己的评论")
	}

	// 删除评论
	err = dao.DeleteCommentById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除评论失败")
	}

	return nil
}
