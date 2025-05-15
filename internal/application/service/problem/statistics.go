package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/problem"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

func Count(params request.QueryProblemParams) (int64, error) {
	query := params2Query(params)
	query.Page = option.NewPagination(0, 0)
	return problem.Query.Count(query)
}
