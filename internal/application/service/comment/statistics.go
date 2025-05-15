package comment

import (
	"STUOJ/internal/domain/comment"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// Count 统计数量
func Count(query querycontext.CommentQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	count, err := comment.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
