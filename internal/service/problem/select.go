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
func SelectById(id uint64, userId uint64, role entity.Role) (model.ProblemData, error) {
	// 获取题目信息
	p, err := dao.SelectProblemById(id)

	if err != nil {
		return model.ProblemData{}, errors.New("获取题目信息失败")
	}

	var testcases []entity.Testcase
	var solutions []entity.Solution

	if p.Status != entity.ProblemPublic {
		if !problemOP(&p, userId, role) {
			return model.ProblemData{}, errors.New("获取失败，没有该题权限")
		}
	}

	if problemOP(&p, userId, role) {
		testcases, _ = dao.SelectTestcasesByProblemId(id)
		solutions, _ = dao.SelectSolutionsByProblemId(id)
	}

	// 封装题目数据
	pd := model.ProblemData{
		Problem:   p,
		Testcases: testcases,
		Solutions: solutions,
	}

	return pd, nil
}

func Select(condition model.ProblemWhere, userId uint64, role entity.Role) (ProblemPage, error) {
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

// 封装题目数据
func wrapProblemDatas(problems []entity.Problem) []model.ProblemData {
	var pds []model.ProblemData

	hideProblemContent(problems)

	for _, p := range problems {
		pd := model.ProblemData{
			Problem: p,
		}

		pds = append(pds, pd)
	}

	return pds
}

// 判断用户是否有该题权限
func problemOP(p *entity.Problem, userId uint64, role entity.Role) bool {
	if role >= entity.RoleAdmin {
		return true
	}
	if role < entity.RoleEditor {
		return false
	}
	for _, i := range p.UserIds {
		if i == userId {
			return true
		}
	}
	for _, i := range p.CollectionUserIds {
		if i == userId {
			return true
		}
	}
	return false
}
