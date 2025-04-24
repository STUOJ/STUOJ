package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/model"
)

// Insert 插入评测点数据
func Insert(req request.CreateTestcaseReq, reqUser model.ReqUser) (int64, error) {
	err := isPermission(reqUser)
	if err != nil {
		return 0, err
	}

	tc1 := testcase.NewTestcase(
		testcase.WithProblemId(req.ProblemId),
		testcase.WithSerial(req.Serial),
		testcase.WithTestInput(req.TestInput),
		testcase.WithTestOutput(req.TestOutput),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(req.ProblemId)
	if err != nil {
		return 0, err
	}

	return tc1.Create()
}
