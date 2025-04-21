package testcase

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

// updateProblemUpdateTime 更新题目更新时间
func updateProblemUpdateTime(id int64) error {
	p1 := problem.NewProblem(
		problem.WithId(id),
	)

	return p1.Update()
}

// isPermission 检查用户权限
func isPermission(reqUser model.ReqUser) error {
	if reqUser.Role < entity.RoleEditor {
		return &errors.ErrUnauthorized
	}
	return nil
}
