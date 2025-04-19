package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/domain/testcase/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"time"
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

	return tc1.Update()
}
