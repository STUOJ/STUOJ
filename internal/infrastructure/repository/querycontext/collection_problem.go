package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go CollectionProblemQueryContext
type CollectionProblemQueryContext struct {
	CollectionId model.FieldList[int64]
	ProblemId    model.FieldList[int64]
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
