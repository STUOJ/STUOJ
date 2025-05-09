package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go TeamSubmissionQueryContext
type TeamSubmissionQueryContext struct {
	TeamId       dto.FieldList[int64]
	SubmissionId dto.FieldList[int64]
	option.QueryParams
	Field field.TeamSubmissionField
}

// applyFilter 应用团队提交记录查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *TeamSubmissionQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.TeamId.Exist() {
		filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.SubmissionId.Exist() {
		filters.Add(field.SubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
