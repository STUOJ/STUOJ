package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type CollectionQueryModel struct {
	Id        model.FieldList[uint64]
	Title     model.Field[string]
	UserId    model.FieldList[uint64]
	Status    model.FieldList[uint64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *CollectionQueryModel) Parse(c *gin.Context) {
	query.Title.Parse(c, "title")
	query.UserId.Parse(c, "user")
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

func (query *CollectionQueryModel) GenerateOptions() *option.QueryOptions {
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
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
