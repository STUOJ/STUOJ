package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	"STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go ProblemTagQueryContext
type ProblemTagQueryContext struct {
	ProblemId dto.Field[int64]
	TagId     dto.Field[int64]
	option.QueryParams
	Field field.ProblemTagField
}

func (query *ProblemTagQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.ProblemId.Exist() {
		filters.Add(field.ProblemTagProblemId, option.OpEqual, query.ProblemId.Value())
	}
	if query.TagId.Exist() {
		filters.Add(field.ProblemTagTagId, option.OpEqual, query.TagId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
