package blog

import (
	"STUOJ/internal/domain/blog"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// Count 统计博客数量
func Count(query querycontext.BlogQueryContext) (int64, error) {
	query.Page = option.NewPagination(0, 0)
	count, err := blog.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
