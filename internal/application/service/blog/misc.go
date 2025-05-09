package blog

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

// 检查权限
func isPermission(b blog.Blog, reqUser request.ReqUser) error {
	if b.UserId.Value() != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
