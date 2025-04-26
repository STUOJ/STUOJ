package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go CollectionUserQueryContext
type CollectionUserQueryContext struct {
	Id           model.FieldList[int64]
	UserId       model.FieldList[int64]
	CollectionId model.FieldList[int64]
	option.QueryParams
	Field field.CollectionUserField
}

// applyFilter 应用收藏用户查询过滤器
func (query *CollectionUserQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.CollectionUserId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.CollectionId.Exist() {
		filters.Add(field.CollectionId, option.OpIn, query.CollectionId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
