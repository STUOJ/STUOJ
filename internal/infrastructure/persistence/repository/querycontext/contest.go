package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go ContestQueryContext
type ContestQueryContext struct {
	Id          dto.FieldList[int64]
	UserId      dto.FieldList[int64]
	Title       dto.Field[string]
	Status      dto.FieldList[entity.ContestStatus]
	Format      dto.FieldList[entity.ContestFormat]
	TeamSize    dto.FieldList[int8]
	StartTime   dto.Field[time.Time]
	EndTime     dto.Field[time.Time]
	BeginStart  dto.Field[time.Time]
	BeginEnd    dto.Field[time.Time]
	FinishStart dto.Field[time.Time]
	FinishEnd   dto.Field[time.Time]
	option.QueryParams
	Field field.ContestField
}

// applyFilter 应用查询过滤器到options
func (query *ContestQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.ContestId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.ContestUserId, option.OpIn, query.UserId.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.ContestTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.ContestStatus, option.OpIn, query.Status.Value())
	}
	if query.Format.Exist() {
		filters.Add(field.ContestFormat, option.OpIn, query.Format.Value())
	}
	if query.TeamSize.Exist() {
		filters.Add(field.ContestTeamSize, option.OpIn, query.TeamSize.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.ContestStartTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.ContestEndTime, option.OpLessEq, query.EndTime.Value())
	}
	if query.BeginStart.Exist() {
		filters.Add(field.ContestStartTime, option.OpGreaterEq, query.BeginStart.Value())
	}
	if query.BeginEnd.Exist() {
		filters.Add(field.ContestStartTime, option.OpLessEq, query.BeginEnd.Value())
	}
	if query.FinishStart.Exist() {
		filters.Add(field.ContestEndTime, option.OpGreaterEq, query.FinishStart.Value())
	}
	if query.FinishEnd.Exist() {
		filters.Add(field.ContestEndTime, option.OpLessEq, query.FinishEnd.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
