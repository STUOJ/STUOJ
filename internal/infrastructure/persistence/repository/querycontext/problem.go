package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go ProblemQueryContext
type ProblemQueryContext struct {
	Id         dto.FieldList[int64]
	Title      dto.Field[string]
	Source     dto.Field[string]
	Status     dto.FieldList[entity.ProblemStatus]
	Difficulty dto.FieldList[entity.Difficulty]
	StartTime  dto.Field[time.Time]
	EndTime    dto.Field[time.Time]
	option2.QueryParams
	Field field.ProblemField
}

// applyFilter 应用查询过滤器到options
func (query *ProblemQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.ProblemId, option2.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.ProblemTitle, option2.OpLike, query.Title.Value())
	}
	if query.Source.Exist() {
		filters.Add(field.ProblemSource, option2.OpLike, query.Source.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.ProblemStatus, option2.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.ProblemCreateTime, option2.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.ProblemCreateTime, option2.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
