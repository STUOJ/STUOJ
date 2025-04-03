package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type CollectionUserQueryModel struct {
	Id           model.FieldList[int64]
	UserId       model.FieldList[int64]
	CollectionId model.FieldList[int64]
	Page         model.QueryPage
	Sort         model.QuerySort
	Field        field.CollectionUserField
}

func (query *CollectionUserQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.CollectionUserId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.CollectionId.Exist() {
		options.Filters.Add(field.CollectionId, option.OpIn, query.CollectionId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
