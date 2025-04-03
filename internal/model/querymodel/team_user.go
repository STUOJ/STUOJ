package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type TeamUserQuery struct {
	TeamId model.FieldList[int64]
	UserId model.FieldList[int64]
	Page   model.QueryPage
	Sort   model.QuerySort
	Field  field.TeamUserField
}

func (query *TeamUserQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.TeamId.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.UserId, option.OpIn, query.UserId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
