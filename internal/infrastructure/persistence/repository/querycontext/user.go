package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go UserQueryContext
type UserQueryContext struct {
	Id        dto.FieldList[int64]
	Username  dto.Field[string]
	Email     dto.Field[string]
	Role      dto.FieldList[entity.Role]
	StartTime dto.Field[time.Time]
	EndTime   dto.Field[time.Time]
	option2.QueryParams
	Field field.UserField
}

// applyFilter 应用查询过滤器到options
func (query *UserQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.UserId, option2.OpIn, query.Id.Value())
	}
	if query.Username.Exist() {
		filters.Add(field.UserUsername, option2.OpLike, query.Username.Value())
	}
	if query.Email.Exist() {
		filters.Add(field.UserEmail, option2.OpEqual, query.Email.Value())
	}
	if query.Role.Exist() {
		filters.Add(field.UserRole, option2.OpIn, query.Role.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.UserCreateTime, option2.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.UserCreateTime, option2.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
