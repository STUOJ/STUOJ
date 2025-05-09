package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option2 "STUOJ/internal/infrastructure/persistence/repository/option"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go CollectionUserQueryContext
type CollectionUserQueryContext struct {
	Id           dto.FieldList[int64]
	UserId       dto.FieldList[int64]
	CollectionId dto.FieldList[int64]
	option2.QueryParams
	Field field.CollectionUserField
}

// applyFilter 应用收藏用户查询过滤器
func (query *CollectionUserQueryContext) applyFilter(options option2.Options) option2.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.CollectionUserId, option2.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.CollectionUserId, option2.OpIn, query.UserId.Value())
	}
	if query.CollectionId.Exist() {
		filters.Add(field.CollectionId, option2.OpIn, query.CollectionId.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
