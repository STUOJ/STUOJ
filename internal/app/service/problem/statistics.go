package problem

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/problem"
)

func GetStatistics(params request.QueryProblemParams) (int64, error) {
	query := params2Query(params)
	return problem.Query.Count(query)
}
