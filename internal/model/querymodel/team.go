package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type TeamQueryModel struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ContestId model.FieldList[int64]
	Name      model.Field[string]
	Status    model.FieldList[int8]
	Page      model.QueryPage
	Sort      model.QuerySort
	Field     field.TeamField
}

func (query *TeamQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.TeamUserId, option.OpIn, query.UserId.Value())
	}
	if query.ContestId.Exist() {
		options.Filters.Add(field.TeamContestId, option.OpIn, query.ContestId.Value())
	}
	if query.Name.Exist() {
		options.Filters.Add(field.TeamName, option.OpLike, query.Name.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.TeamStatus, option.OpIn, query.Status.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
