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
func SelectById(id uint64, admin ...bool) (model.ProblemData, error) {
	// 获取题目信息
	p, err := dao.SelectProblemById(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题目信息失败")
	}

	// 获取题目标签
	tags, err := dao.SelectTagsByProblemId(id)
	if err != nil {
		return model.ProblemData{}, errors.New("获取题目标签失败")
	}

	var testcases []entity.Testcase
	var solutions []entity.Solution

	if len(admin) > 0 && admin[0] {

		// 获取评测点数据
		testcases, err = dao.SelectTestcasesByProblemId(id)
		if err != nil {
			return model.ProblemData{}, errors.New("获取评测点数据失败")
		}

		// 获取题解数据
		solutions, err = dao.SelectSolutionsByProblemId(id)
		if err != nil {
			return model.ProblemData{}, errors.New("获取题解数据失败")
		}
	} else {
		if p.Status != entity.ProblemStatusPublic {
			return model.ProblemData{}, errors.New("题目未公开")
		}
	}

	// 封装题目数据
	pd := model.ProblemData{
		Problem:   p,
		Tags:      tags,
		Testcases: testcases,
		Solutions: solutions,
	}

	return pd, nil
}

func Select(condition dao.ProblemWhere) (ProblemPage, error) {
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
