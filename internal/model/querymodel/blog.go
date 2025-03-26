package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogQueryModel struct {
	Id        model.FieldList[uint64]
	UserId    model.FieldList[uint64]
	ProblemId model.FieldList[uint64]
	Title     model.Field[string]
	Status    model.FieldList[uint64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *BlogQueryModel) Parse(c *gin.Context) {
	query.Title.Parse(c, "title")
	query.Status.Parse(c, "status")
	query.ProblemId.Parse(c, "problem")
	query.UserId.Parse(c, "user")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *BlogQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.BlogId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.BlogUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.BlogProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.BlogTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.BlogStatus, option.OpIn, query.Status.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.BlogCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.BlogCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
