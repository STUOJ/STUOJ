package testcase

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/model"
)

// Delete 根据ID删除评测点数据
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
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
	err = updateProblemUpdateTime(tc0.ProblemId)
	if err != nil {
		return err
	}

	return tc0.Delete()
}
