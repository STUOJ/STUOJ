package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go TestcaseQueryModel
type TestcaseQueryModel struct {
	Id        model.FieldList[int64]
	ProblemId model.FieldList[int64]
	Page      option.Pagination
	Sort      option.Sort
	Field     field.TestcaseField
}

func (query *TestcaseQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.TestcaseId, option.OpIn, query.Id.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.TestcaseProblemId, option.OpIn, query.ProblemId.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
