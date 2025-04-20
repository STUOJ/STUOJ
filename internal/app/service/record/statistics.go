package record

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/submission"
)

// Count 统计提交记录数量
func Count(req request.QuerySubmissionParams) (int64, error) {
	query := params2Query(req)
	count, err := submission.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
