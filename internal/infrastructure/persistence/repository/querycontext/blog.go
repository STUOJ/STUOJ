package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go BlogQueryContext
type BlogQueryContext struct {
	Id        dto.FieldList[int64]
	UserId    dto.FieldList[int64]
	ProblemId dto.FieldList[int64]
	Title     dto.Field[string]
	Status    dto.FieldList[entity.BlogStatus]
	StartTime dto.Field[time.Time]
	EndTime   dto.Field[time.Time]
	option.QueryParams
	Field field.BlogField
}

func (query *BlogQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.BlogId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.BlogUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.BlogProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.BlogTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.BlogStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.BlogCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.BlogCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
