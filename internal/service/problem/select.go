package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/solution"
	"errors"
	"slices"
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

	op := problemOP(p, userId, role)

	if p.Status < entity.ProblemPublic && !op {
		return entity.Problem{}, errors.New("无权限")
	}

	if where.Solutions.Exist() && where.Solutions.Value() && op {
		solutions, _ := solution.SelectByProblemId(id)
		p.Solutions = solutions
	}

	if where.Testcases.Exist() && where.Testcases.Value() && op {
		testcases, _ := dao.SelectTestcasesByProblemId(id)
		p.Testcases = testcases
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

// 判断用户是否有该题权限
func problemOP(p entity.Problem, userId uint64, role entity.Role) bool {
	if role >= entity.RoleAdmin {
		return true
	}
	if role < entity.RoleEditor {
		return false
	}
	if slices.Contains(p.UserIds, userId) {
		return true
	}
	return slices.Contains(p.CollectionUserIds, userId)
}
