package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type Score struct {
	shared.Valueobject[int64]
}

func NewScore(value int64) Score {
	var s Score
	s.Set(value)
	return s
}

func (s Score) Verify() error {
	if s.Value() < 0 {
		return errors.New("分数不合法")
	}
	return nil
}
