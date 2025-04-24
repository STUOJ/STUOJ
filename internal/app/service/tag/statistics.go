package tag

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/tag"
)

// Count 统计用户数量
func Count(req request.QueryTagParams) (int64, error) {
	query := params2Query(req)
	count, err := tag.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
