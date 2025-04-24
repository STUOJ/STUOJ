package testcase

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/testcase"
)

func GetStatistics(params request.QueryTestcaseParams) (int64, error) {
	query := params2Query(params)
	return testcase.Query.Count(query)
}
