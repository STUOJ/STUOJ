package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"
)

//go:generate go run ../../../utils/gen/querymodel_gen.go TeamSubmissionQuery
type TeamSubmissionQuery struct {
	TeamId       model.FieldList[int64]
	SubmissionId model.FieldList[int64]
	Page         option.Pagination
	Sort         option.Sort
	Field        field.TeamSubmissionField
}

func (query *TeamSubmissionQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.TeamId.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.SubmissionId.Exist() {
		options.Filters.Add(field.SubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
