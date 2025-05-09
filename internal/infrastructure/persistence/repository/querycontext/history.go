package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go HistoryQueryContext
type HistoryQueryContext struct {
	Id         dto.FieldList[int64]
	UserId     dto.FieldList[int64]
	ProblemId  dto.FieldList[int64]
	Title      dto.Field[string]
	Difficulty dto.FieldList[int64]
	StartTime  dto.Field[time.Time]
	EndTime    dto.Field[time.Time]
	Operation  dto.FieldList[entity.Operation]
	option.QueryParams
	Field field.HistoryField
}

// applyFilter 应用历史记录查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *HistoryQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.HistoryId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.HistoryUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.HistoryProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.HistoryTitle, option.OpLike, query.Title.Value())
	}
	if query.Difficulty.Exist() {
		filters.Add(field.HistoryDifficulty, option.OpIn, query.Difficulty.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.HistoryCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.HistoryCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	if query.Operation.Exist() {
		filters.Add(field.HistoryOperation, option.OpIn, query.Operation.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
