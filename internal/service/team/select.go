package team

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"fmt"
	"log"
	"slices"
)

type TeamPage struct {
	Teams []entity.Team `json:"teams"`
	model.Page
}

// 根据ID查询团队
func SelectById(id uint64, userId uint64, role entity.Role) (entity.Team, error) {
	// 获取团队信息
	tm, err := dao.SelectTeamById(id)

	if err != nil {
		return entity.Team{}, errors.New("获取团队失败")
	}
	flag := false
	if tm.Status != entity.TeamPublic && role < entity.RoleAdmin && tm.UserId != userId {
		if slices.Contains(tm.UserIds, userId) {
			flag = true
		}
	} else {
		flag = true
	}
	if !flag {
		return entity.Team{}, errors.New("没有权限查看该团队")
	}

	uCondition := model.UserWhere{}
	uCondition.Id.Set(tm.UserIds)
	uCondition.Page.Set(uint64(1))
	uCondition.Size.Set(uint64(100))
	uCondition.OrderBy.Set(fmt.Sprintf("FIELD(tbl_user.id,%s)", utils.Uint64SliceToString(tm.UserIds)))
	tm.Users, err = dao.SelectUsers(uCondition)
	if err != nil {
		return tm, errors.New("获取用户信息失败")
	}

	return tm, nil
}

// 查询团队
func Select(condition model.TeamWhere, uid uint64, role entity.Role) (TeamPage, error) {
	if !condition.Status.Exist() {
		condition.Status.Set([]uint64{uint64(entity.TeamPublic)})
	} else {
		for _, v := range condition.Status.Value() {
			if entity.TeamStatus(v) < entity.TeamPublic {
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
	users, err := dao.SelectTeams(condition)
	if err != nil {
		log.Println(err)
		return TeamPage{}, errors.New("查询团队失败")
	}

	count, err := dao.CountTeams(condition)
	if err != nil {
		log.Println(err)
		return TeamPage{}, errors.New("查询统计失败")
	}
	uPage := TeamPage{
		Teams: users,
		Page: model.Page{
			Total: count,
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
		},
	}

	return uPage, nil
}
