package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go TagQueryContext
type TagQueryContext struct {
	Id   dto.FieldList[int64]
	Name dto.FieldList[string]
	option2.QueryParams
	Field field.TagField
}

// applyFilter 应用标签查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *TagQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.TagId, option2.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.TagName, option2.OpLike, query.Name.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
