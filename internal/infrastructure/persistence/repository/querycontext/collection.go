package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go CollectionQueryContext
type CollectionQueryContext struct {
	Id        dto.FieldList[int64]
	Title     dto.Field[string]
	UserId    dto.FieldList[int64]
	Status    dto.FieldList[entity.CollectionStatus]
	StartTime dto.Field[time.Time]
	EndTime   dto.Field[time.Time]
	option.QueryParams
	Field field.CollectionField
}

// applyFilter 应用收藏查询过滤器
func (query *CollectionQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.CollectionId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		filters.Add(field.CollectionTitle, option.OpLike, query.Title.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.CollectionStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.CollectionCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.CollectionCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
