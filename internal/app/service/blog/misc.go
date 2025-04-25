package blog

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// 检查权限
func isPermission(b blog.Blog, reqUser model.ReqUser) error {
	if b.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
