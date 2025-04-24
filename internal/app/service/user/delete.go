package user

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// Delete 根据Id删除用户
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	if reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}

	u1 := user.NewUser(
		user.WithId(id),
	)

	return u1.Delete()
}
