package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go CollectionProblemQueryModel
type CollectionProblemQueryModel struct {
	CollectionId model.FieldList[int64]
	ProblemId    model.FieldList[int64]
	Page         option.Pagination
	Sort         option.Sort
	Field        field.CollectionProblemField
}

func (query *CollectionProblemQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.CollectionId.Exist() {
		options.Filters.Add(field.CollectionProblemCollectionId, option.OpIn, query.CollectionId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.CollectionProblemProblemId, option.OpIn, query.ProblemId.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
