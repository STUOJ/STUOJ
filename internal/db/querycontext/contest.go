package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
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

func (query *ContestQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.ContestId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.ContestUserId, option.OpIn, query.UserId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.ContestTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.ContestStatus, option.OpIn, query.Status.Value())
	}
	if query.Format.Exist() {
		options.Filters.Add(field.ContestFormat, option.OpIn, query.Format.Value())
	}
	if query.TeamSize.Exist() {
		options.Filters.Add(field.ContestTeamSize, option.OpIn, query.TeamSize.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpLessEq, query.EndTime.Value())
	}
	if query.BeginStart.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpGreaterEq, query.BeginStart.Value())
	}
	if query.BeginEnd.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpLessEq, query.BeginEnd.Value())
	}
	if query.FinishStart.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpGreaterEq, query.FinishStart.Value())
	}
	if query.FinishEnd.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpLessEq, query.FinishEnd.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
