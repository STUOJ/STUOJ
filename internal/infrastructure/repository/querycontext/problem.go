package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go ProblemQueryContext
type ProblemQueryContext struct {
	Id         model.FieldList[int64]
	Title      model.Field[string]
	Source     model.Field[string]
	Status     model.FieldList[uint8]
	Difficulty model.FieldList[uint8]
	StartTime  model.Field[time.Time]
	EndTime    model.Field[time.Time]
	option.QueryParams
	Field field.ProblemField
}

// applyFilter 应用查询过滤器到options
func (query *ProblemQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.ProblemId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.ProblemTitle, option.OpLike, query.Title.Value())
	}
	if query.Source.Exist() {
		filters.Add(field.ProblemSource, option.OpLike, query.Source.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.ProblemStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.ProblemCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.ProblemCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
