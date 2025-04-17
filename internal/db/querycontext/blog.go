package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go BlogQueryContext
type BlogQueryContext struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ProblemId model.FieldList[int64]
	Title     model.Field[string]
	Status    model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	option.QueryParams
	Field field.BlogField
}

func (query *BlogQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.BlogId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.BlogUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.BlogProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.BlogTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.BlogStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.BlogCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.BlogCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
