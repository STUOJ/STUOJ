package collection

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/collection"
	"STUOJ/internal/model"
)

// DeleteLogic 逻辑删除
func DeleteLogic(id int64, reqUser model.ReqUser) error {
	// 查询
	qc := querycontext.CollectionQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUserId()
	cl0, _, err := collection.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(cl0, reqUser)
	if err != nil {
		return err
	}

	// 逻辑删除
	cl1 := collection.NewCollection(
		collection.WithId(id),
		collection.WithStatus(entity.CollectionDeleted),
	)

	return cl1.Update()
}

// Delete 根据Id删除题单
func Delete(id int64, reqUser model.ReqUser) error {
	// 查询题单
	queryContext := querycontext.CollectionQueryContext{}
	queryContext.Field.SelectId().SelectUserId()
	queryContext.Id.Add(id)
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
