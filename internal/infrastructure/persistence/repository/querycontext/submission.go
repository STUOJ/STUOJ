package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go SubmissionQueryContext
type SubmissionQueryContext struct {
	Id        dto.FieldList[int64]
	UserId    dto.FieldList[int64]
	ProblemId dto.FieldList[int64]
	Status    dto.FieldList[entity.JudgeStatus]
	Language  dto.FieldList[int64]
	StartTime dto.Field[time.Time]
	EndTime   dto.Field[time.Time]
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
