package testcase

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

// updateProblemUpdateTime 更新题目更新时间
func updateProblemUpdateTime(id int64) error {
	p1 := problem.NewProblem(
		problem.WithId(id),
	)

	return p1.Update()
}

// isPermission 检查用户权限
func isPermission(reqUser request.ReqUser) error {
	if reqUser.Role < entity.RoleEditor {
		return &errors.ErrUnauthorized
	}
	return nil
}
