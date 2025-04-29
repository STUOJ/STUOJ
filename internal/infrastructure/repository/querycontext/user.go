package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go UserQueryContext
type UserQueryContext struct {
	Id        model.FieldList[int64]
	Username  model.Field[string]
	Email     model.Field[string]
	Role      model.FieldList[uint8]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	option.QueryParams
	Field field.UserField
}

// applyFilter 应用查询过滤器到options
func (query *UserQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.UserId, option.OpIn, query.Id.Value())
	}
	if query.Username.Exist() {
		filters.Add(field.UserUsername, option.OpLike, query.Username.Value())
	}
	if query.Email.Exist() {
		filters.Add(field.UserEmail, option.OpEqual, query.Email.Value())
	}
	if query.Role.Exist() {
		filters.Add(field.UserRole, option.OpIn, query.Role.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.UserCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.UserCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
