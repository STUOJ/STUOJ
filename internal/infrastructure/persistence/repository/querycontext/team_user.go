package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go TeamUserQueryContext
type TeamUserQueryContext struct {
	TeamId dto.FieldList[int64]
	UserId dto.FieldList[int64]
	option2.QueryParams
	Field field.TeamUserField
}

// applyFilter 应用团队成员查询过滤器
func (query *TeamUserQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.TeamId.Exist() {
		filters.Add(field.TeamId, option2.OpIn, query.TeamId.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.UserId, option2.OpIn, query.UserId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
