package querycontext

import (
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
	"time"
)

//go:generate go run ../../../../dev/gen/querycontext_gen.go CollectionQueryContext
type CollectionQueryContext struct {
	Id        model.FieldList[int64]
	Title     model.Field[string]
	UserId    model.FieldList[int64]
	Status    model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
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
