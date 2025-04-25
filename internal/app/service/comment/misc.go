package comment

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// 检查权限
func isPermission(cm comment.Comment, reqUser model.ReqUser) error {
	if cm.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
