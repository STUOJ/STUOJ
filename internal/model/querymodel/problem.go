package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type ProblemQueryModel struct {
	Id        model.FieldList[uint64]
	Title     model.Field[string]
	Source    model.Field[string]
	Status    model.FieldList[uint8]
	Tag       model.FieldList[uint64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *ProblemQueryModel) Parse(c *gin.Context) {
	query.Title.Parse(c, "title")
	query.Source.Parse(c, "source")
	query.Status.Parse(c, "status")
	query.Tag.Parse(c, "tag")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *ProblemQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.ProblemId, option.OpIn, query.Id.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.ProblemTitle, option.OpLike, query.Title.Value())
	}
	if query.Source.Exist() {
		options.Filters.Add(field.ProblemSource, option.OpLike, query.Source.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.ProblemStatus, option.OpIn, query.Status.Value())
	}
	if query.Tag.Exist() {
		options.Filters.Add(field.ProblemTag, option.OpHave, query.Tag.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.ProblemCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
