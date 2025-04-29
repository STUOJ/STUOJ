package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go ContestQueryContext
type ContestQueryContext struct {
	Id          model.FieldList[int64]
	UserId      model.FieldList[int64]
	Title       model.Field[string]
	Status      model.FieldList[int8]
	Format      model.FieldList[int8]
	TeamSize    model.FieldList[int8]
	StartTime   model.Field[time.Time]
	EndTime     model.Field[time.Time]
	BeginStart  model.Field[time.Time]
	BeginEnd    model.Field[time.Time]
	FinishStart model.Field[time.Time]
	FinishEnd   model.Field[time.Time]
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
