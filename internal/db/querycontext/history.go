package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go HistoryQueryContext
type HistoryQueryContext struct {
	Id         model.FieldList[int64]
	UserId     model.FieldList[int64]
	ProblemId  model.FieldList[int64]
	Title      model.Field[string]
	Difficulty model.FieldList[int64]
	StartTime  model.Field[time.Time]
	EndTime    model.Field[time.Time]
	Operation  model.FieldList[int8]
	option.QueryParams
	Field field.HistoryField
}

func (query *HistoryQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.HistoryId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.HistoryUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.HistoryProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.HistoryTitle, option.OpLike, query.Title.Value())
	}
	if query.Difficulty.Exist() {
		options.Filters.Add(field.HistoryDifficulty, option.OpIn, query.Difficulty.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.HistoryCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.HistoryCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	if query.Operation.Exist() {
		options.Filters.Add(field.HistoryOperation, option.OpIn, query.Operation.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
