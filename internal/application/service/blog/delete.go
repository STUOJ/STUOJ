package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
)

// DeleteLogic 逻辑删除博客
func DeleteLogic(id int64, reqUser request.ReqUser) error {
	// 查询
	qc := querycontext.BlogQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUserId()
	b0, _, err := blog.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(b0, reqUser)
	if err != nil {
		return err
	}

	// 逻辑删除
	b1 := blog.NewBlog(
		blog.WithId(id),
		blog.WithStatus(entity.BlogDeleted),
	)

	return b1.Update()
}

// Delete 根据Id删除博客
func Delete(id int64, reqUser request.ReqUser) error {
	// 查询
	qc := querycontext.BlogQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUserId()
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
