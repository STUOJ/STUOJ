package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go TeamUserQueryContext
type TeamUserQueryContext struct {
	TeamId dto.FieldList[int64]
	UserId dto.FieldList[int64]
	option.QueryParams
	Field field.TeamUserField
}

// applyFilter 应用团队成员查询过滤器
func (query *TeamUserQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.TeamId.Exist() {
		filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.UserId, option.OpIn, query.UserId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
