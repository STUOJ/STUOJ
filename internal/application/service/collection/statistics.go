package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/collection"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

func Count(req request.QueryCollectionParams) (int64, error) {
	query := params2Model(req)
	query.Page = option.NewPagination(0, 0)
	return collection.Query.Count(query)
}
