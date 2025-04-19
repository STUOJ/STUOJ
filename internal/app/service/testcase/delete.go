package testcase

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"time"
)

// Delete 根据ID删除评测点数据
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	if reqUser.Role < entity.RoleEditor {
		return &errors.ErrUnauthorized
	}

	// 查询
	qc := querycontext.TestcaseQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	tc0, _, err := testcase.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 更新题目更新时间
	pqc := querycontext.ProblemQueryContext{}
	pqc.Id.Add(tc0.ProblemId)
	pqc.Field.SelectId().SelectUpdateTime()
	p0, _, err := problem.Query.SelectOne(pqc)
	if err != nil {
		return errors.ErrNotFound.WithMessage("找不到对应的题目")
	}
	p0.UpdateTime = time.Now()
	err = p0.Update()
	if err != nil {
		return errors.ErrInternalServer.WithMessage("更新题目更新时间失败")
	}

	return tc0.Delete()
}
