package contest

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新比赛
func Update(ct entity.Contest, uid uint64, role entity.Role) error {
	// 查询比赛
	c0, err := dao.SelectContestById(ct.Id)
	if err != nil {
		log.Println(err)
		return errors.New("比赛不存在")
	}

	if role < entity.RoleAdmin && c0.UserId != uid {
		return errors.New("没有权限，只能操作自己的比赛")

	}

	// 更新比赛
	updateTime := time.Now()
	c0.Status = ct.Status
	c0.Format = ct.Format
	c0.TeamSize = ct.TeamSize
	c0.StartTime = ct.StartTime
	c0.EndTime = ct.EndTime
	c0.UpdateTime = updateTime

	// 更新比赛
	err = dao.UpdateContestById(c0)
	if err != nil {
		log.Println(err)
		return errors.New("修改比赛失败")
	}

	return nil
}
