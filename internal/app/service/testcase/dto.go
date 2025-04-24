package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/testcase"
	"STUOJ/utils"
)

func params2Query(params request.QueryTestcaseParams) (query querycontext.TestcaseQueryContext) {
	if params.Problem != nil {
		if ids, err := utils.StringToInt64Slice(*params.Problem); err == nil {
			query.ProblemId.Set(ids)
		}
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	return query
}

func domain2Resp(dm testcase.Testcase) response.TestcaseData {
	return response.TestcaseData{
		Id:         dm.Id,
		ProblemId:  dm.ProblemId,
		Serial:     dm.Serial,
		TestInput:  dm.TestInput.String(),
		TestOutput: dm.TestOutput.String(),
	}
}
