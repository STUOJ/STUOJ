package blog

import (
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
)

// 根据ID删除博客（检查用户ID）
func Delete(id uint64, reqUser model.ReqUser) error {
	queryContext := querycontext.BlogQueryContext{}
	blog, _, err := blog.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	err = isPermission(blog, reqUser)
	if err != nil {
		return err
	}
	return blog.Delete()
}
