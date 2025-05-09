package user

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// 检查权限
func isAdminPermission(reqUser model.ReqUser) error {
	if reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}

// 检查权限
func isPermission(id int64, reqUser model.ReqUser) error {
	if reqUser.Id != id && reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}
	return nil
}
