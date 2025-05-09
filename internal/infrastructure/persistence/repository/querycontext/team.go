package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go TeamQueryContext
type TeamQueryContext struct {
	Id        dto.FieldList[int64]
	UserId    dto.FieldList[int64]
	ContestId dto.FieldList[int64]
	Name      dto.Field[string]
	Status    dto.FieldList[entity.TeamStatus]
	option2.QueryParams
	Field field.TeamField
}

// applyFilter 应用团队查询过滤器
func (query *TeamQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.TeamId, option2.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.TeamUserId, option2.OpIn, query.UserId.Value())
	}
	if query.ContestId.Exist() {
		filters.Add(field.TeamContestId, option2.OpIn, query.ContestId.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.TeamName, option2.OpLike, query.Name.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.TeamStatus, option2.OpIn, query.Status.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
