package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go CollectionProblemQueryContext
type CollectionProblemQueryContext struct {
	CollectionId dto.FieldList[int64]
	ProblemId    dto.FieldList[int64]
	option.QueryParams
	Field field.CollectionProblemField
}

func (query *CollectionProblemQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.CollectionId.Exist() {
		filters.Add(field.CollectionProblemCollectionId, option.OpIn, query.CollectionId.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.CollectionProblemProblemId, option.OpIn, query.ProblemId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
