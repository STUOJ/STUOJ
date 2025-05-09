package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go SolutionQueryContext
type SolutionQueryContext struct {
	Id         dto.FieldList[int64]
	ProblemId  dto.FieldList[int64]
	LanguageId dto.FieldList[int64]
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
