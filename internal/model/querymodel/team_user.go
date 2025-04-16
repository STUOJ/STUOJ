package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go TeamUserQuery
type TeamUserQuery struct {
	TeamId model.FieldList[int64]
	UserId model.FieldList[int64]
	Page   option.Pagination
	Sort   option.Sort
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
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
