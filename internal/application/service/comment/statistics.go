package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/comment"
)

// Count 统计数量
func Count(req request.QueryCommentParams) (int64, error) {
	query := params2Query(req)
	count, err := comment.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
