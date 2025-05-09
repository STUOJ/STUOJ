package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Status struct {
	model.Valueobject[entity.ProblemStatus]
}

func NewStatus(value entity.ProblemStatus) Status {
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
