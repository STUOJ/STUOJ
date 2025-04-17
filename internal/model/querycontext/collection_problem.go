package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go CollectionProblemQueryContext
type CollectionProblemQueryContext struct {
	CollectionId model.FieldList[int64]
	ProblemId    model.FieldList[int64]
	option.QueryParams
	Field field.CollectionProblemField
}

func (query *CollectionProblemQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.CollectionId.Exist() {
		options.Filters.Add(field.CollectionProblemCollectionId, option.OpIn, query.CollectionId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.CollectionProblemProblemId, option.OpIn, query.ProblemId.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
