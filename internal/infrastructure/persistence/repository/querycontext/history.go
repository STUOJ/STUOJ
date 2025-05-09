package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
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
	option2.QueryParams
	Field field.HistoryField
}

// applyFilter 应用历史记录查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *HistoryQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.HistoryId, option2.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.HistoryUserId, option2.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.HistoryProblemId, option2.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.HistoryTitle, option2.OpLike, query.Title.Value())
	}
	if query.Difficulty.Exist() {
		filters.Add(field.HistoryDifficulty, option2.OpIn, query.Difficulty.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.HistoryCreateTime, option2.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.HistoryCreateTime, option2.OpLessEq, query.EndTime.Value())
	}
	if query.Operation.Exist() {
		filters.Add(field.HistoryOperation, option2.OpIn, query.Operation.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
