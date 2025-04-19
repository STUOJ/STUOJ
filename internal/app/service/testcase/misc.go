package testcase

import (
	"STUOJ/internal/db/entity"
	"STUOJ/internal/errors"
	"STUOJ/internal/model"
)

func isPermission(reqUser model.ReqUser) error {
	if reqUser.Role < entity.RoleEditor {
		return &errors.ErrUnauthorized
	}
	return nil
}
