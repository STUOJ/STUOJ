package problem

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/model"
)

// DeleteLogic 逻辑删除
func DeleteLogic(id int64, reqUser model.ReqUser) error {
	// 查询
	qc := querycontext.ProblemQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectAll()
	_, problemMap, err := problem.Query.SelectOne(qc, problem.QueryUser())
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(problemMap, reqUser)
	if err != nil {
		return err
	}

	// 逻辑删除
	p1 := problem.NewProblem(
		problem.WithId(id),
		problem.WithStatus(entity.ProblemDeleted),
	)

	return p1.Update()
}
