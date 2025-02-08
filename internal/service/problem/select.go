package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
)

type ProblemPage struct {
	Problems []entity.Problem `json:"problems"`
	model.Page
}

// 根据ID查询题目数据
func SelectById(id uint64, userId uint64, role entity.Role, where model.ProblemWhere) (entity.Problem, error) {
	where.ScoreUserId.Set(userId)
	p, err := dao.SelectProblemById(id, where)

	if err != nil {
		return entity.Problem{}, errors.New("获取题目信息失败")
	}

	if p.Status != entity.ProblemPublic && role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range p.UserIds {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[userId]; !exists {
			return entity.Problem{}, errors.New("没有该题权限")
		}
	}

	return p, nil
}

func Select(condition model.ProblemWhere, userId uint64, role entity.Role) (ProblemPage, error) {
	condition.ScoreUserId.Set(userId)
	if !condition.Status.Exist() {
		condition.Status.Set([]uint64{uint64(entity.ProblemPublic)})
	} else {
		for _, v := range condition.Status.Value() {
			if entity.ProblemStatus(v) < entity.ProblemPublic {
				if role <= entity.RoleUser {
					condition.Status.Set(entity.ProblemPublic)
				} else if role == entity.RoleEditor {
					condition.UserId.Set(userId)
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
	problems, err := dao.SelectProblems(condition)
	if err != nil {
		return ProblemPage{}, errors.New("获取题目信息失败")
	}

	hideProblemContent(problems)

	count, err := dao.CountProblems(condition)
	if err != nil {
		return ProblemPage{}, errors.New("获取题目总数失败")
	}

	pPage := ProblemPage{
		Problems: problems,
		Page: model.Page{
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
			Total: count,
		},
	}

	return pPage, nil
}

func hideProblemContent(problems []entity.Problem) {
	for i := range problems {
		problems[i].Description = ""
		problems[i].Input = ""
		problems[i].Output = ""
		problems[i].SampleInput = ""
		problems[i].SampleOutput = ""
		problems[i].Hint = ""
	}
}
