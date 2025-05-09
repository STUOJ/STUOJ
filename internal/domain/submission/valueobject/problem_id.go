package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type ProblemId struct {
	model.Valueobject[int64]
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
