package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Role struct {
	model.Valueobject[entity.Role]
}

func NewRole(value entity.Role) Role {
	var s Role
	s.Set(value)
	return s
}

func (s Role) Verify() error {
	if !s.Value().IsValid() {
		return errors.ErrStatus
	}
	return nil
}
