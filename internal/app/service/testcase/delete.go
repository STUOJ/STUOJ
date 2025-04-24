package testcase

import (
	"STUOJ/internal/domain/testcase"
	"STUOJ/internal/model"
)

// Delete 根据ID删除评测点数据
func Delete(id int64, reqUser model.ReqUser) error {
	// 检查权限
	err := isPermission(reqUser)
	if err != nil {
		return err
	}

	tc1 := testcase.NewTestcase(
		testcase.WithId(id),
	)

	// 更新题目更新时间
	err = updateProblemUpdateTime(id)
	if err != nil {
		return err
	}

	return tc1.Delete()
}
