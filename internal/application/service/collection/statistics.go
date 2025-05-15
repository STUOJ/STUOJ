package collection

import (
	"STUOJ/internal/domain/collection"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

func Count(query querycontext.CollectionQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	return collection.Query.Count(query)
}
