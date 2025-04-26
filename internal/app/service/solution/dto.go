package solution

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/app/dto/response"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/model/option"
	"STUOJ/utils"
)

func params2Query(params request.QuerySolutionParams) (query querycontext.SolutionQueryContext) {
	if params.Language != nil {
		if ids, err := utils.StringToInt64Slice(*params.Language); err == nil {
			query.LanguageId.Set(ids)
		}
	}
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

func domain2Resp(dm solution.Solution) response.SolutionData {
	return response.SolutionData{
		Id:         dm.Id,
		ProblemId:  dm.ProblemId,
		LanguageId: dm.LanguageId,
		SourceCode: dm.SourceCode.String(),
	}
}
