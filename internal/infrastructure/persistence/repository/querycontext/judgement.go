package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go JudgementQueryContext
type JudgementQueryContext struct {
	Id           dto.FieldList[int64]
	SubmissionId dto.FieldList[int64]
	TestcaseId   dto.FieldList[int64]
	Status       dto.FieldList[entity.JudgeStatus]
	option.QueryParams
	Field field.JudgementField
}

// applyFilter 应用评测记录查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *JudgementQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.JudgementId, option.OpIn, query.Id.Value())
	}
	if query.SubmissionId.Exist() {
		filters.Add(field.JudgementSubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	if query.TestcaseId.Exist() {
		filters.Add(field.JudgementTestcaseId, option.OpIn, query.TestcaseId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.JudgementStatus, option.OpIn, query.Status.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
