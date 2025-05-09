package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

type Operation struct {
	shared.Valueobject[entity.Operation]
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
