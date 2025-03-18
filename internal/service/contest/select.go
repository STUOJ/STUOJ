package contest

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

type ContestPage struct {
	Contests []entity.Contest `json:"contests"`
	model.Page
}

// 根据ID查询比赛
func SelectById(id uint64, userId uint64, role entity.Role) (entity.Contest, error) {
	// 获取比赛信息
	ct, err := dao.SelectContestById(id)

	if err != nil {
		return entity.Contest{}, errors.New("获取比赛失败")
	}

	if ct.Status < entity.ContestReady && role < entity.RoleAdmin && ct.UserId != userId {
		return entity.Contest{}, errors.New("没有权限查看该比赛")
	}

	return ct, nil
}

// 查询比赛
func Select(condition model.ContestWhere, uid uint64, role entity.Role) (ContestPage, error) {
	if !condition.Status.Exist() {
		condition.Status.Set([]uint64{uint64(entity.ContestReady)})
	} else {
		for _, v := range condition.Status.Value() {
			if entity.ContestStatus(v) < entity.ContestReady {
				if role < entity.RoleAdmin {
					condition.UserId.Set([]uint64{uid})
				}
			}
		}
	}
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	users, err := dao.SelectContests(condition)
	if err != nil {
		log.Println(err)
		return ContestPage{}, errors.New("查询比赛失败")
	}

	count, err := dao.CountContests(condition)
	if err != nil {
		log.Println(err)
		return ContestPage{}, errors.New("查询统计失败")
	}
	uPage := ContestPage{
		Contests: users,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}
