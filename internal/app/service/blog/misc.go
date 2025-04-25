package blog

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/blog"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

func isPermission(b blog.Blog, reqUser model.ReqUser) error {
	// 检查权限
	if b.UserId != reqUser.Id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
