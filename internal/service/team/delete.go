package team

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据ID删除团队
func Delete(id uint64, uid uint64, role entity.Role) error {
	// 查询团队
	t0, err := dao.SelectBlogById(id)
	if err != nil {
		log.Println(err)
		return errors.New("团队不存在")
	}

	if role < entity.RoleAdmin && t0.UserId != uid {
		return errors.New("没有权限，只能删除自己的团队")
	}

	// 删除团队
	err = dao.DeleteTeamById(id)
	if err != nil {
		log.Println(err)
		return errors.New("删除团队失败")
	}

	return nil
}

func DeleteUser(tmid uint64, delUid uint64, uid uint64, role entity.Role) error {
	tu := entity.TeamUser{
		TeamId: tmid,
		UserId: delUid,
	}
	_, err := dao.SelectUserById(delUid)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}
	t0, err := dao.SelectTeamById(tmid)
	if err != nil {
		log.Println(err)
		return errors.New("团队不存在")
	}
	if role < entity.RoleAdmin {
		if t0.UserId != uid {
			return errors.New("没有权限，只能操作自己的团队")
		}
	}
	if t0.UserId == delUid {
		return errors.New("团队创建者不能删除")
	}
	for _, i := range t0.UserIds {
		if i == delUid {
			err = dao.DeleteTeamUser(tu)
			if err != nil {
				log.Println(err)
				return errors.New("删除用户失败")
			}
			return nil
		}
	}
	return errors.New("用户不在团队中")
}
