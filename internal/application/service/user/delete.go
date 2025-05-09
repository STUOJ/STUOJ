package user

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/user"
)

// Delete 根据Id删除用户
func Delete(id int64, reqUser request.ReqUser) error {
	// 检查权限
	err := isAdminPermission(reqUser)
	if err != nil {
		return err
	}

	u1 := user.NewUser(
		user.WithId(id),
	)

	return u1.Delete()
}
