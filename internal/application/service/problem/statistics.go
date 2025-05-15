package problem

import (
	"STUOJ/internal/domain/problem"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

func Count(query querycontext.ProblemQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	return problem.Query.Count(query)
}
