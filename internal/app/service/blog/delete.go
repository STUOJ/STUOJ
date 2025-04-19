package blog

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/model"
)

// Delete 根据Id删除博客
func Delete(id uint64, reqUser model.ReqUser) error {
	// 查询
	qc := querycontext.BlogQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	b0, _, err := blog.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(b0, reqUser)
	if err != nil {
		return err
	}

	return b0.Delete()
}
