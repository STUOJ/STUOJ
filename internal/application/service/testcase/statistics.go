package testcase

import (
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

func Count(query querycontext.TestcaseQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	return testcase.Query.Count(query)
}
