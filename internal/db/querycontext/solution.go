package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go SolutionQueryContext
type SolutionQueryContext struct {
	Id         model.FieldList[int64]
	ProblemId  model.FieldList[int64]
	LanguageId model.FieldList[int64]
	option.QueryParams
	Field field.SolutionField
}

// applyFilter 应用查询过滤器到options
func (query *SolutionQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.SolutionId, option.OpIn, query.Id.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.SolutionProblemId, option.OpIn, query.ProblemId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
