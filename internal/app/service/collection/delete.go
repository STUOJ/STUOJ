package collection

import (
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/model"
	"STUOJ/internal/model/querycontext"
)

// 根据ID删除题单
func Delete(id uint64, reqUser model.ReqUser) error {
	// 查询题单
	queryContext := querycontext.CollectionQueryContext{}
	queryContext.Id.Add(reqUser.ID)
	c0, _, err := collection.Query.SelectOne(queryContext)
	if err != nil {
		return err
	}
	// 检查权限
	err = isPermission(c0, reqUser)
	if err != nil {
		return err
	}
	// 删除题单
	return c0.Delete()
}
