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

	if p.Status != entity.ProblemPublic && role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range p.UserIds {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[userId]; !exists {
			return model.ProblemData{}, errors.New("没有该题权限")
		}
	}

	if role >= entity.RoleEditor {
		testcases, err = dao.SelectTestcasesByProblemId(id)
		solutions, err = dao.SelectSolutionsByProblemId(id)
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
	problems, err := dao.SelectProblem(condition)
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
