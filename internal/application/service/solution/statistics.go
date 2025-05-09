package solution

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/solution"
)

func GetStatistics(params request.QuerySolutionParams) (int64, error) {
	query := params2Query(params)
	return solution.Query.Count(query)
}
