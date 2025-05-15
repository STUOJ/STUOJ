package record

import (
	"STUOJ/internal/domain/submission"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// Count 统计提交记录数量
func Count(query querycontext.SubmissionQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	count, err := submission.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
