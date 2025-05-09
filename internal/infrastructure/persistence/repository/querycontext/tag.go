package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go TagQueryContext
type TagQueryContext struct {
	Id   dto.FieldList[int64]
	Name dto.FieldList[string]
	option.QueryParams
	Field field.TagField
}

// applyFilter 应用标签查询过滤条件
// 根据查询参数设置过滤条件，并返回更新后的options对象
func (query *TagQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.TagId, option.OpIn, query.Id.Value())
	}
	if query.Name.Exist() {
		filters.Add(field.TagName, option.OpLike, query.Name.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
