package comment

import (
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model"
)

// DeleteLogic 逻辑删除
func DeleteLogic(id int64, reqUser model.ReqUser) error {
	// 查询
	qc := querycontext.CommentQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId().SelectUserId()
	cm0, _, err := comment.Query.SelectOne(qc)
	if err != nil {
		return err
	}

	// 检查权限
	err = isPermission(cm0, reqUser)
	if err != nil {
		return err
	}

	// 逻辑删除
	cm1 := comment.NewComment(
		comment.WithId(id),
		comment.WithStatus(entity.CommentDeleted),
	)

	return cm1.Update()
}
