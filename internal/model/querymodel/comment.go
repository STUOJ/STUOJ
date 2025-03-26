package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentQueryModel struct {
	Id        model.FieldList[uint64]
	UserId    model.FieldList[uint64]
	BlogId    model.Field[uint64]
	Status    model.FieldList[uint64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *CommentQueryModel) Parse(c *gin.Context) {
	query.UserId.Parse(c, "user")
	query.BlogId.Parse(c, "blog")
	query.Status.Parse(c, "status")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
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
	return options
}
