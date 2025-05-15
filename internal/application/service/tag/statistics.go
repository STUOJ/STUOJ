package tag

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/infrastructure/persistence/repository/option"
)

// Count 统计用户数量
func Count(req request.QueryTagParams) (int64, error) {
	query := params2Query(req)
	query.Page = option.NewPagination(0, 0)
	count, err := tag.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
