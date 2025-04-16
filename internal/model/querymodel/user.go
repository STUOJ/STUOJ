package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go UserQueryModel
type UserQueryModel struct {
	Id        model.FieldList[int64]
	Username  model.Field[string]
	Role      model.FieldList[int8]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      option.Pagination
	Sort      option.Sort
	Field     field.UserField
}

func (query *UserQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.UserId, option.OpIn, query.Id.Value())
	}
	if query.Username.Exist() {
		options.Filters.Add(field.UserUsername, option.OpLike, query.Username.Value())
	}
	if query.Role.Exist() {
		options.Filters.Add(field.UserRole, option.OpIn, query.Role.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.UserCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.UserCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
