package team

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入团队
func Insert(tm entity.Team) (uint64, error) {
	var err error

	updateTime := time.Now()
	tm.UpdateTime = updateTime
	tm.CreateTime = updateTime

	// 插入团队
	tm.Id, err = dao.InsertTeam(tm)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败")
	}

	return tm.Id, nil
}

func InsertUser(tmid uint64, addUid uint64, uid uint64, role entity.Role) error {
	tu := entity.TeamUser{
		TeamId: tmid,
		UserId: addUid,
	}

	_, err := dao.SelectUserById(addUid)
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
	if t0.UserId == addUid {
		return errors.New("不能添加自己为团队成员")
	}
	for _, i := range t0.UserIds {
		if i == addUid {
			return errors.New("用户已存在团队中")
		}
	}
	err = dao.InsertTeamUser(tu)
	if err != nil {
		log.Println(err)
		return errors.New("添加用户失败")
	}
	return nil
}
