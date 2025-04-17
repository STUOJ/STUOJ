package blog

import (
	"STUOJ/internal/app/dto/request"
	"STUOJ/internal/domain/blog"
)

// 统计博客数量
func GetStatistics(req request.QueryBlogParams) (int64, error) {
	query := params2Model(req)
	return blog.Query.Count(query)
}
