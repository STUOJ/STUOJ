package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Time struct {
	model.Valueobject[int64]
}

func NewTime(value int64) Time {
	var t Time
	t.Set(value)
	return t
}

func (t Time) Verify() error {
	if t.Value() <= 0 {
		return errors.New("time must be positive")
	}
	return nil
}
