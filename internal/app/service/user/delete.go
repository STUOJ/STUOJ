package user

import (
	"STUOJ/internal/db/querycontext"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/model"
	"log"
)

// Delete 根据Id删除用户
func Delete(id uint64, reqUser model.ReqUser) error {
	// 查询用户
	qc := querycontext.UserQueryContext{}
	u0, _, err := user.Query.SelectOne(qc)
	if err != nil {
		log.Println(err)
	}

	// 检查权限
	err = isPermission(reqUser)
	if err != nil {
		log.Println(err)
	}

	return u0.Delete()
}
