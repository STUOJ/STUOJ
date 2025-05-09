package submission

import (
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/pkg/errors"
)

func (_Query) SelectACUsers(pid, size int64) ([]map[string]any, error) {
	res, err := dao.SubmissionStore.SelectACUsers(pid, size)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage("查询AC用户失败")
	}
	return res, nil
}
