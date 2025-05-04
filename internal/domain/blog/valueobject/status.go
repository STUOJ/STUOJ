package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"fmt"
)

type Status struct {
	model.Valueobject[entity.BlogStatus]
}

func NewStatus(value entity.BlogStatus) Status {
	var s Status
	s.Set(value)
	return s
}

func (s Status) Verify() error {
	if !s.Value().IsValid() {
		return ErrStatus
	}
	return nil
}

var ErrStatus = fmt.Errorf("status error")
