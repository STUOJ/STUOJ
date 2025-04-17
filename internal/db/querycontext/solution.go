package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go SolutionQueryContext
type SolutionQueryContext struct {
	Id         model.FieldList[int64]
	ProblemId  model.FieldList[int64]
	LanguageID model.FieldList[int64]
	option.QueryParams
	Field field.SolutionField
}

func (query *SolutionQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.SolutionId, option.OpIn, query.Id.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.SolutionProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.LanguageID.Exist() {
		options.Filters.Add(field.SolutionLanguageId, option.OpIn, query.LanguageID.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
