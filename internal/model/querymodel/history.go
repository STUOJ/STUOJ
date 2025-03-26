package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type HistoryQueryModel struct {
	Id         model.FieldList[uint64]
	UserId     model.FieldList[uint64]
	ProblemId  model.FieldList[uint64]
	Title      model.Field[string]
	Difficulty model.FieldList[uint64]
	StartTime  model.Field[time.Time]
	EndTime    model.Field[time.Time]
	Page       model.QueryPage
	Sort       model.QuerySort
}

func (query *HistoryQueryModel) Parse(c *gin.Context) {
	query.UserId.Parse(c, "user")
	query.ProblemId.Parse(c, "problem")
	query.Title.Parse(c, "title")
	query.Difficulty.Parse(c, "difficulty")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *HistoryQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.HistoryId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.HistoryUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.HistoryProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.HistoryTitle, option.OpLike, query.Title.Value())
	}
	if query.Difficulty.Exist() {
		options.Filters.Add(field.HistoryDifficulty, option.OpIn, query.Difficulty.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.HistoryCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.HistoryCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
