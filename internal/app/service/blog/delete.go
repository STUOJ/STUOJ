package blog

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/model"
)

// 根据Id删除博客（检查用户Id）
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
