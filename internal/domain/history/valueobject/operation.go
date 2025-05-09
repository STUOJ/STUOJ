package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Operation struct {
	model.Valueobject[entity.Operation]
}

func NewOperation(value entity.Operation) Operation {
	var s Operation
	s.Set(value)
	return s
}

func (s Operation) Verify() error {
	if !s.Value().IsValid() {
		return errors.ErrStatus
	}
	return nil
}
