package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go CollectionUserQueryContext
type CollectionUserQueryContext struct {
	Id           dto.FieldList[int64]
	UserId       dto.FieldList[int64]
	CollectionId dto.FieldList[int64]
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
