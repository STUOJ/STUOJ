package testcase

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/infrastructure/persistence/repository/option"
)

func Count(params request.QueryTestcaseParams) (int64, error) {
	query := params2Query(params)
	query.Page = option.NewPagination(0, 0)
	return testcase.Query.Count(query)
}
