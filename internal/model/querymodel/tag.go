package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go TagQueryModel
type TagQueryModel struct {
	Id    model.FieldList[int64]
	Name  model.FieldList[string]
	Page  option.Pagination
	Sort  option.Sort
	Field field.TagField
}

func (query *TagQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.TagId, option.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		options.Filters.Add(field.TagName, option.OpLike, query.Name.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
