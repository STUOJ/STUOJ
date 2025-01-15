package user

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据ID删除用户
func DeleteById(id uint64) error {
	// 查询用户
	_, err := dao.SelectUserById(id)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}

	// 删除用户
	err = dao.DeleteUserById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除用户失败")
	}

	return nil
}
