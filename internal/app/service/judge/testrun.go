package judge

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query"
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/runner"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
)

func TestRun(req request.TestRunReq, reqUser model.ReqUser) (response.TestRunRes, error) {
	languageQuery := querycontext.LanguageQueryContext{}
	languageQuery.Id.Add(req.LanguageID)
	languageQuery.Field = *query.LanguageAllField
	languageDomain, _, err := language.Query.SelectOne(languageQuery)
	if err != nil {
		return response.TestRunRes{}, err
	}
	languageMapId := languageDomain.MapId
	submission := runner.RunnerSubmission{
		SourceCode: req.SourceCode,
		LanguageId: uint64(languageMapId),
	}
	testcaseSubmission := runner.RunnerTestcaseSubmission{
		Input: req.Stdin,
	}
	submission.Testcase = append(submission.Testcase, testcaseSubmission)
	runResult := runner.Runner.CodeRun(submission)
	res := response.Result2TestRunRes(runResult)
	return res, nil
}
