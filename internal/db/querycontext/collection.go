package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go CollectionQueryContext
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

func (query *CollectionQueryContext) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.CollectionId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.CollectionTitle, option.OpLike, query.Title.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.CollectionStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.CollectionCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.CollectionCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
