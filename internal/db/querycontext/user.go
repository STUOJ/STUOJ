package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go UserQueryContext
type UserQueryContext struct {
	Id        model.FieldList[uint64]
	Username  model.Field[string]
	Role      model.FieldList[uint8]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	option.QueryParams
	Field field.UserField
}

func (query *UserQueryContext) GenerateOptions() *option.QueryOptions {
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
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
