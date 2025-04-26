package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go TeamUserQuery
type TeamUserQuery struct {
	TeamId model.FieldList[int64]
	UserId model.FieldList[int64]
	option.QueryParams
	Field field.TeamUserField
}

// applyFilter 应用团队成员查询过滤器
func (query *TeamUserQuery) applyFilter(options option.Options) option.Options {
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
