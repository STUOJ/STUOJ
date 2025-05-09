package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go SubmissionQueryContext
type SubmissionQueryContext struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ProblemId model.FieldList[int64]
	Status    model.FieldList[entity.JudgeStatus]
	Language  model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	option.QueryParams
	Field field.SubmissionField
}

// applyFilter 应用查询过滤器到options
func (query *SubmissionQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.SubmissionId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.SubmissionUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		filters.Add(field.SubmissionProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.SubmissionStatus, option.OpIn, query.Status.Value())
	}
	if query.Language.Exist() {
		filters.Add(field.SubmissionLanguageId, option.OpIn, query.Language.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.SubmissionCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.SubmissionCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
