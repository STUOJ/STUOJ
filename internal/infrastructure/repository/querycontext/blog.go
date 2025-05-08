package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go BlogQueryContext
type BlogQueryContext struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ProblemId model.FieldList[int64]
	Title     model.Field[string]
	Status    model.FieldList[entity.BlogStatus]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
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
