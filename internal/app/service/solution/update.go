package solution

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/model"
)

// Update 根据ID更新题解
func Update(req request.UpdateSolutionReq, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	// 查询题解
	qc := querycontext.SolutionQueryContext{}
	qc.Id.Add(req.Id)
	qc.Field.SelectAll()
	s0, _, err := solution.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	s1 := solution.NewSolution(
		solution.WithId(s0.Id),
		solution.WithProblemId(s0.ProblemId),
		solution.WithLanguageId(req.LanguageId),
		solution.WithSourceCode(req.SourceCode),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(req.ProblemId)
	if err != nil {
		return err
	}

	return s1.Update()
}
