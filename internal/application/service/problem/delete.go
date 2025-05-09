package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// DeleteLogic 逻辑删除
func DeleteLogic(id int64, reqUser request.ReqUser) error {
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
