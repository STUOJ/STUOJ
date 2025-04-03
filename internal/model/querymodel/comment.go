package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

type CommentQueryModel struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	BlogId    model.Field[int64]
	Status    model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
	Field     field.CommentField
}

func (query *CommentQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.CommentId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.CommentUserId, option.OpIn, query.UserId.Value())
	}
	if query.BlogId.Exist() {
		options.Filters.Add(field.CommentBlogId, option.OpEqual, query.BlogId.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.CommentStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.CommentCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.CommentCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
