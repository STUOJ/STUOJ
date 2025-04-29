package testcase

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/testcase"
)

func GetStatistics(params request.QueryTestcaseParams) (int64, error) {
	query := params2Query(params)
	return testcase.Query.Count(query)
}
