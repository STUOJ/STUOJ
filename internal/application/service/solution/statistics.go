package solution

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/infrastructure/persistence/repository/option"
)

func Count(params request.QuerySolutionParams) (int64, error) {
	query := params2Query(params)
	query.Page = option.NewPagination(0, 0)
	return solution.Query.Count(query)
}
