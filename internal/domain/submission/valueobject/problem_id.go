package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/pkg/errors"
)

type ProblemId struct {
	shared.Valueobject[int64]
}

func NewProblemId(value int64) ProblemId {
	var p ProblemId
	p.Set(value)
	return p
}

func (p ProblemId) Verify() error {
	if p.Value() < 0 {
		return errors.ErrProblemId
	}
	return nil
}
