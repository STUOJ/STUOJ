package tag

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/tag"
	"STUOJ/internal/model"
)

// Delete 根据ID删除标签
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	// 查询
	qc := querycontext.TagQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectAll()
	t0, _, err := tag.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	return t0.Delete()
}
