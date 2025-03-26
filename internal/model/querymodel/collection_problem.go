package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type CollectionProblemQueryModel struct {
	CollectionId model.FieldList[uint64]
	ProblemId    model.FieldList[uint64]
	Page         model.QueryPage
	Sort         model.QuerySort
}

func (query *CollectionProblemQueryModel) Parse(c *gin.Context) {
	query.CollectionId.Parse(c, "collection")
	query.ProblemId.Parse(c, "problem")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *CollectionProblemQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.CollectionId.Exist() {
		options.Filters.Add(field.CollectionProblemCollectionId, option.OpIn, query.CollectionId.Value())
	}
	if query.ProblemId.Exist() {
		options.Filters.Add(field.CollectionProblemProblemId, option.OpIn, query.ProblemId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
