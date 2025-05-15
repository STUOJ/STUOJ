package history

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/history"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
)

// Count 统计数量
func Count(req request.QueryHistoryParams) (int64, error) {
	query := params2Query(req)
	query.Page = option.NewPagination(0, 0)
	count, err := history.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
