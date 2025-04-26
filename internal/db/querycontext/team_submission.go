package querycontext

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/model"
	"STUOJ/internal/model/option"
)

//go:generate go run ../../../utils/gen/querycontext_gen.go TeamSubmissionQuery
type TeamSubmissionQuery struct {
	TeamId       model.FieldList[int64]
	SubmissionId model.FieldList[int64]
	option.QueryParams
	Field field.TeamSubmissionField
}

func (query *TeamSubmissionQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.TeamId.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.SubmissionId.Exist() {
		options.Filters.Add(field.SubmissionId, option.OpIn, query.SubmissionId.Value())
	}
	options.Filters.AddFiter(query.ExtraFilters.Conditions...)
	options.Page = query.Page
	options.Sort = query.Sort
	options.Field = &query.Field
	return options
}
