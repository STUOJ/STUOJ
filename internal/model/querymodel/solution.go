package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type SolutionQueryModel struct {
	Id         model.FieldList[uint64]
	ProblemId  model.FieldList[uint64]
	LanguageID model.FieldList[uint64]
	Page       model.QueryPage
	Sort       model.QuerySort
}

func (query *SolutionQueryModel) Parse(c *gin.Context) {
	query.ProblemId.Parse(c, "problem")
	query.LanguageID.Parse(c, "language")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *SolutionQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.SolutionId, option.OpIn, query.Id.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.SolutionProblemId, option.OpIn, query.ProblemId.Value())
	}
	if query.LanguageID.Exist() {
		options.Filters.Add(field.SolutionLanguageId, option.OpIn, query.LanguageID.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
