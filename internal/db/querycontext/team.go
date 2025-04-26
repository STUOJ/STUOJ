package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go TeamQueryContext
type TeamQueryContext struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ContestId model.FieldList[int64]
	Name      model.Field[string]
	Status    model.FieldList[int8]
	option.QueryParams
	Field field.TeamField
}

// applyFilter 应用团队查询过滤器
func (query *TeamQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.TeamId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.TeamUserId, option.OpIn, query.UserId.Value())
	}
	if query.ContestId.Exist() {
		filters.Add(field.TeamContestId, option.OpIn, query.ContestId.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.TeamName, option.OpLike, query.Name.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.TeamStatus, option.OpIn, query.Status.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
