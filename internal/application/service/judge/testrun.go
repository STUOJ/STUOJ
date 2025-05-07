package judge

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/runner"
	"STUOJ/internal/infrastructure/repository/query"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
)

func TestRun(req request.TestRunReq, reqUser model.ReqUser) (response.TestRunRes, error) {
	languageQuery := querycontext.LanguageQueryContext{}
	languageQuery.Id.Add(req.LanguageId)
	languageQuery.Field = *query.LanguageAllField
	languageDomain, _, err := language.Query.SelectOne(languageQuery)
	if err != nil {
		return response.TestRunRes{}, err
	}
	languageMapId := languageDomain.MapId.Value()
	submission := runner.RunnerSubmission{
		SourceCode: req.SourceCode,
		LanguageId: int64(languageMapId),
	}
	testcaseSubmission := runner.RunnerTestcaseSubmission{
		Input: req.Stdin,
	}
	submission.Testcase = append(submission.Testcase, testcaseSubmission)
	runResult := runner.Runner.CodeRun(submission)
	res := response.Result2TestRunRes(runResult)
	return res, nil
}
