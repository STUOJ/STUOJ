package comment

import (
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// 检查权限
func isPermission(cm comment.Comment, reqUser model.ReqUser) error {
	if cm.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
