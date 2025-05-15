package history

import (
	"STUOJ/internal/domain/history"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// Count 统计数量
func Count(query querycontext.HistoryQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	count, err := history.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
