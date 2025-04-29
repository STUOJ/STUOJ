package querycontext

import (
	field2 "STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go TeamSubmissionQuery
type TeamSubmissionQuery struct {
	TeamId       model.FieldList[int64]
	SubmissionId model.FieldList[int64]
	option.QueryParams
	Field field2.TeamSubmissionField
}

// applyFilter 应用团队提交记录查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *TeamSubmissionQuery) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.TeamId.Exist() {
		filters.Add(field2.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.SubmissionId.Exist() {
		filters.Add(field2.SubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
