package querycontext

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/entity/field"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"time"
)

//go:generate go run ../../../../../dev/gen/querycontext_gen.go CommentQueryContext
type CommentQueryContext struct {
	Id        dto.FieldList[int64]
	UserId    dto.FieldList[int64]
	BlogId    dto.Field[int64]
	Status    dto.FieldList[entity.CommentStatus]
	StartTime dto.Field[time.Time]
	EndTime   dto.Field[time.Time]
	option.QueryParams
	Field field.CommentField
}

// applyFilter 应用查询过滤器到options
func (query *CommentQueryContext) applyFilter(options option.Options) option.Options {
	filters := options.GetFilters()
	if query.Id.Exist() {
		filters.Add(field.CommentId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		filters.Add(field.CommentUserId, option.OpIn, query.UserId.Value())
	}
	if query.BlogId.Exist() {
		filters.Add(field.CommentBlogId, option.OpEqual, query.BlogId.Value())
	}
	if query.Status.Exist() {
		filters.Add(field.CommentStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		filters.Add(field.CommentCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		filters.Add(field.CommentCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	filters.AddFiter(query.ExtraFilters.Conditions...)
	return options
}
