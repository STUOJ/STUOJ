package solution

import (
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

func Count(query querycontext.SolutionQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	return solution.Query.Count(query)
}
