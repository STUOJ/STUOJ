package judge

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/runner"
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
	languageMapId := languageDomain.MapId
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
