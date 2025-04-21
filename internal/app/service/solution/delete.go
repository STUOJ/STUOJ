package solution

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/model"
)

// Delete 根据ID删除题解
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	// 查询
	qc := querycontext.SolutionQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	s0, _, err := solution.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 更新题目更新时间
	err = updateProblemUpdateTime(s0.ProblemId)
	if err != nil {
		return err
	}

	return s0.Delete()
}
