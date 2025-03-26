package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type TeamUserQuery struct {
	TeamId model.FieldList[uint64]
	UserId model.FieldList[uint64]
	Page   model.QueryPage
	Sort   model.QuerySort
}

func (query *TeamUserQuery) Parse(c *gin.Context) {
	query.TeamId.Parse(c, "team")
	query.UserId.Parse(c, "user")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *TeamUserQuery) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.TeamId.Exist() {
		options.Filters.Add(field.TeamId, option.OpIn, query.TeamId.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.UserId, option.OpIn, query.UserId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
