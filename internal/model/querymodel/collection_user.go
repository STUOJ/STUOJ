package querymodel

import (
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/model"

	"github.com/gin-gonic/gin"
)

type CollectionUserQueryModel struct {
	Id           model.FieldList[uint64]
	UserId       model.FieldList[uint64]
	CollectionId model.FieldList[uint64]
	Page         model.QueryPage
	Sort         model.QuerySort
}

func (query *CollectionUserQueryModel) Parse(c *gin.Context) {
	query.UserId.Parse(c, "user")
	query.CollectionId.Parse(c, "collection")
	query.Page.Parse(c)
	query.Sort.Parse(c)
}

func (query *CollectionUserQueryModel) GenerateOptions() *option.QueryOptions {
	options := option.NewQueryOptions()
	if query.Id.Exist() {
		options.Filters.Add(field.CollectionUserId, option.OpIn, query.Id.Value())
	}
	if query.UserId.Exist() {
		options.Filters.Add(field.CollectionUserId, option.OpIn, query.UserId.Value())
	}
	if query.CollectionId.Exist() {
		options.Filters.Add(field.CollectionId, option.OpIn, query.CollectionId.Value())
	}
	query.Page.InsertOptions(options)
	query.Sort.InsertOptions(options)
	return options
}
