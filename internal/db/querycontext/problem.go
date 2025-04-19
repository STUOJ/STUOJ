package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go ProblemQueryContext
type ProblemQueryContext struct {
	Id        model.FieldList[uint64]
	Title     model.Field[string]
	Source    model.Field[string]
	Status    model.FieldList[uint8]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	option.QueryParams
	Field field.ProblemField
}

func (query *ProblemQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.ProblemId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.ProblemTitle, option.OpLike, query.Title.Value())
	}
	if query.Source.Exist() {
		options.Filters.Add(field.ProblemSource, option.OpLike, query.Source.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.ProblemStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
