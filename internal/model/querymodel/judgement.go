package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type JudgementQueryModel struct {
	Id           model.FieldList[int64]
	SubmissionId model.FieldList[int64]
	TestcaseId   model.FieldList[int64]
	Status       model.FieldList[int64]
	Page         model.QueryPage
	Sort         model.QuerySort
	Field        field.JudgementField
}

func (query *JudgementQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.JudgementId, option.OpIn, query.Id.Value())
	}
	if query.SubmissionId.Exist() {
		options.Filters.Add(field.JudgementSubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	if query.TestcaseId.Exist() {
		options.Filters.Add(field.JudgementTestcaseId, option.OpIn, query.TestcaseId.Value())
	}
	if query.Status.Exist() {
		options.Filters.Add(field.JudgementStatus, option.OpIn, query.Status.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	options.Field = &query.Field
	return options
}
