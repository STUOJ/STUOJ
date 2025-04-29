package blog

import (
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// 检查权限
func isPermission(b blog.Blog, reqUser model.ReqUser) error {
	if b.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
