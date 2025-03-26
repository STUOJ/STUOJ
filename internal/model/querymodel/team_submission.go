package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type TeamSubmissionQuery struct {
	TeamId       model.FieldList[uint64]
	SubmissionId model.FieldList[uint64]
	Page         model.QueryPage
	Sort         model.QuerySort
}

func (query *TeamSubmissionQuery) Parse(c *gin.Context) {
	query.TeamId.Parse(c, "team")
	query.SubmissionId.Parse(c, "submission")
	query.Page.Parse(c)
	query.Sort.Parse(c)
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
