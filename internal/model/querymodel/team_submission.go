package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

type TeamSubmissionQuery struct {
	TeamId       model.FieldList[uint64]
	SubmissionId model.FieldList[uint64]
	Page         model.QueryPage
	Sort         model.QuerySort
}

func (query *TeamSubmissionQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.TeamId.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.SubmissionId.Exist() {
		options.Filters.Add(field.SubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
