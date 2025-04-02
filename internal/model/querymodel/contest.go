package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type ContestQueryModel struct {
	Id          model.FieldList[uint64]
	UserId      model.FieldList[uint64]
	Title       model.Field[string]
	Status      model.FieldList[uint8]
	Format      model.FieldList[uint8]
	TeamSize    model.FieldList[uint8]
	StartTime   model.Field[time.Time]
	EndTime     model.Field[time.Time]
	BeginStart  model.Field[time.Time]
	BeginEnd    model.Field[time.Time]
	FinishStart model.Field[time.Time]
	FinishEnd   model.Field[time.Time]
	Page        model.QueryPage
	Sort        model.QuerySort
}

func (query *ContestQueryModel) Parse(c *gin.Context) {
	query.UserId.Parse(c, "user")
	query.Status.Parse(c, "status")
	query.Title.Parse(c, "title")
	query.Format.Parse(c, "format")
	query.TeamSize.Parse(c, "team_size")
	query.StartTime.Parse(c, "start_time")
	query.EndTime.Parse(c, "end_time")
	query.BeginStart.Parse(c, "begin_start")
	query.BeginEnd.Parse(c, "begin_end")
	query.FinishStart.Parse(c, "finish_start")
	query.FinishEnd.Parse(c, "finish_end")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *ContestQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.ContestId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.ContestUserId, option.OpIn, query.UserId.Value())
	}
	if query.Title.Exist() {
		options.Filters.Add(field.ContestTitle, option.OpLike, query.Title.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.ContestStatus, option.OpIn, query.Status.Value())
	}
	if query.Format.Exist() {
		options.Filters.Add(field.ContestFormat, option.OpIn, query.Format.Value())
	}
	if query.TeamSize.Exist() {
		options.Filters.Add(field.ContestTeamSize, option.OpIn, query.TeamSize.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpLessEq, query.EndTime.Value())
	}
	if query.BeginStart.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpGreaterEq, query.BeginStart.Value())
	}
	if query.BeginEnd.Exist() {
		options.Filters.Add(field.ContestStartTime, option.OpLessEq, query.BeginEnd.Value())
	}
	if query.FinishStart.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpGreaterEq, query.FinishStart.Value())
	}
	if query.FinishEnd.Exist() {
		options.Filters.Add(field.ContestEndTime, option.OpLessEq, query.FinishEnd.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
