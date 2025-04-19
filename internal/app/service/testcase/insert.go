package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

func Insert(req request.CreateTestcaseReq, reqUser model.ReqUser) (uint64, error) {
	t := testcase.NewTestcase(
		testcase.WithProblemId(req.ProblemId),
		testcase.WithSerial(req.Serial),
		testcase.WithTestInput(req.TestInput),
		testcase.WithTestOutput(req.TestOutput),
	)

	// 检查权限
	if reqUser.Role < entity.RoleEditor {
		return 0, &errors.ErrUnauthorized
	}

	return t.Create()
}
