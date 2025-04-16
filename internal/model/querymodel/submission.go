package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
	"time"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go SubmissionQueryModel
type SubmissionQueryModel struct {
	Id        model.FieldList[int64]
	UserId    model.FieldList[int64]
	ProblemId model.FieldList[int64]
	Status    model.FieldList[int8]
	Language  model.FieldList[int64]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      option.Pagination
	Sort      option.Sort
	Field     field.SubmissionField
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
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
