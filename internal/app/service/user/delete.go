package user

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
	"log"
)

// Delete 根据Id删除用户
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	if reqUser.Role < entity.RoleAdmin {
		return &errors.ErrUnauthorized
	}

	// 查询用户
	qc := querycontext.UserQueryContext{}
	qc.Id.Add(id)
	qc.Field.SelectId()
	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		log.Println(err)
	}

	return u0.Delete()
}
