package comment

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/comment"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

// 检查权限
func isPermission(cm comment.Comment, reqUser request.ReqUser) error {
	if cm.UserId.Value() != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
