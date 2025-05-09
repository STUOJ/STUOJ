package judge

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/runner"
	"STUOJ/internal/model"
)

func TestRun(req request.TestRunReq, reqUser model.ReqUser) (response.TestRunRes, error) {
	languageMapId, err := SelectLanguageMapId(req.LanguageId)
	if err != nil {
		return response.TestRunRes{}, err
	}
	submission := runner.RunnerSubmission{
		SourceCode: req.SourceCode,
		LanguageId: int64(languageMapId),
		Testcase:   make([]runner.RunnerTestcaseSubmission, 1),
	}
	testcaseSubmission := runner.RunnerTestcaseSubmission{
		Input: req.Stdin,
	}
	submission.Testcase[0] = testcaseSubmission
	runResult := runner.Runner.CodeRun(submission)
	res := response.Result2TestRunRes(runResult)
	return res, nil
}
