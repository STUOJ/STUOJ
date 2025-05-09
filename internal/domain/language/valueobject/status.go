package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

type Status struct {
	shared.Valueobject[entity.LanguageStatus]
}

func NewStatus(value entity.LanguageStatus) Status {
	var s Status
	s.Set(value)
	return s
}

func (s Status) Verify() error {
	if !s.Value().IsValid() {
		return errors.ErrStatus
	}
	return nil
}
