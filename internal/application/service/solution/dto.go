package solution

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
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
		Id:         dm.Id.Value(),
		ProblemId:  dm.ProblemId.Value(),
		LanguageId: dm.LanguageId.Value(),
		SourceCode: dm.SourceCode.Value(),
	}
}
