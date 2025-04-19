package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/testcase"
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
	err := isPermission(reqUser)
	if err != nil {
		return 0, err
	}

	return t.Create()
}
