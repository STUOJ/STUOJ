package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Status struct {
	model.Valueobject[entity.CommentStatus]
}

func NewStatus(value entity.CommentStatus) Status {
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
