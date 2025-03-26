package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

type SubmissionQueryModel struct {
	Id        model.FieldList[uint64]
	UserId    model.FieldList[uint64]
	ProblemId model.FieldList[uint64]
	Status    model.FieldList[uint8]
	Language  model.FieldList[uint64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *SubmissionQueryModel) Parse(c *gin.Context) {
	query.UserId.Parse(c, "user")
	query.ProblemId.Parse(c, "problem")
	query.Status.Parse(c, "status")
	query.Language.Parse(c, "language")
	timePreiod := &model.Period{}
	err := timePreiod.GetPeriod(c)
	if err == nil {
		query.StartTime.Set(timePreiod.StartTime)
		query.EndTime.Set(timePreiod.EndTime)
	}
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *SubmissionQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.SubmissionId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.SubmissionUserId, option.OpIn, query.UserId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.SubmissionProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.SubmissionStatus, option.OpIn, query.Status.Value())
	}
	if query.Language.Exist() {
		options.Filters.Add(field.SubmissionLanguageId, option.OpIn, query.Language.Value())
	}
	if query.StartTime.Exist() {
		options.Filters.Add(field.SubmissionCreateTime, option.OpGreaterEq, query.StartTime.Value())
	}
	if query.EndTime.Exist() {
		options.Filters.Add(field.SubmissionCreateTime, option.OpLessEq, query.EndTime.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
