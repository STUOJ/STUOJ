package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/domain/testcase/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// Update 根据ID更新评测点数据
func Update(req request.UpdateTestcaseReq, reqUser model.ReqUser) error {
	// 检查权限
	if reqUser.Role < entity.RoleEditor {
		return &errors.ErrUnauthorized
	}

	// 查询
	qc := querycontext.TestcaseQueryContext{}
	qc.Id.Add(req.Id)
	qc.Field.SelectAll()
	tc0, _, err := testcase.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	tc0.Serial = req.Serial
	tc0.TestInput = valueobject.NewTestInput(req.TestInput)
	tc0.TestOutput = valueobject.NewTestOutput(req.TestOutput)

	tc1 := testcase.NewTestcase(
		testcase.WithId(tc0.Id),
		testcase.WithSerial(req.Serial),
		testcase.WithTestInput(req.TestInput),
		testcase.WithTestOutput(req.TestOutput),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(tc0.ProblemId)
	if err != nil {
		return err
	}

	return tc1.Update()
}
