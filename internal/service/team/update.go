package team

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新团队
func Update(tm entity.Team, uid uint64, role entity.Role) error {
	// 查询团队
	t0, err := dao.SelectTeamById(tm.Id)
	if err != nil {
		log.Println(err)
		return errors.New("团队不存在")
	}

	if role < entity.RoleAdmin && t0.UserId != uid {
		return errors.New("没有权限，只能操作自己的团队")

	}

	// 更新团队
	updateTime := time.Now()
	t0.Title = tm.Title
	t0.Description = tm.Description
	t0.Status = tm.Status
	t0.UpdateTime = updateTime

	// 更新团队
	err = dao.UpdateTeamById(t0)
	if err != nil {
		log.Println(err)
		return errors.New("修改团队失败")
	}

	return nil
}
