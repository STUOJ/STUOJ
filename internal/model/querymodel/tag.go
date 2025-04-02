package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type TagQueryModel struct {
	Id   model.FieldList[uint64]
	Name model.Field[string]
	Page model.QueryPage
	Sort model.QuerySort
}

func (query *TagQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.TagId, option.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		options.Filters.Add(field.TagName, option.OpLike, query.Name.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
