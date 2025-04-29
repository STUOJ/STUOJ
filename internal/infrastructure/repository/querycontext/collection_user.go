package querycontext

import (
	field2 "STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go CollectionUserQueryContext
type CollectionUserQueryContext struct {
	Id           model.FieldList[int64]
	UserId       model.FieldList[int64]
	CollectionId model.FieldList[int64]
	option.QueryParams
	Field field2.CollectionUserField
}

// applyFilter 应用收藏用户查询过滤器
func (query *CollectionUserQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field2.CollectionUserId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field2.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.CollectionId.Exist() {
		filters.Add(field2.CollectionId, option.OpIn, query.CollectionId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
