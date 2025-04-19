package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/blog"
)

// Count 统计博客数量
func Count(req request.QueryBlogParams) (int64, error) {
	query := params2Query(req)
	count, err := blog.Query.Count(query)
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
