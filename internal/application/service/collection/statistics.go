package collection

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/collection"
)

func GetStatistics(req request.QueryCollectionParams) (int64, error) {
	query := params2Model(req)
	return collection.Query.Count(query)
}
