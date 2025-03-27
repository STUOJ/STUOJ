package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type TestcaseQueryModel struct {
	Id        model.FieldList[uint64]
	ProblemId model.FieldList[uint64]
	Page      model.QueryPage
	Sort      model.QuerySort
}

func (query *TestcaseQueryModel) Parse(c *gin.Context) {
	query.ProblemId.Parse(c, "problem")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *TestcaseQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.TestcaseId, option.OpIn, query.Id.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.TestcaseProblemId, option.OpIn, query.ProblemId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
