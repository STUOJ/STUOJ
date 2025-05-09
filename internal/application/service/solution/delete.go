package solution

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/domain/solution"
)

// Delete 根据ID删除题解
func Delete(id int64, reqUser request.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	s1 := solution.NewSolution(
		solution.WithId(id),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(id)
	if err != nil {
		return err
	}

	return s1.Delete()
}
