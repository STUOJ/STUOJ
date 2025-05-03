package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
)

type Source struct {
	model.Valueobject[string]
}

func NewSource(value string) Source {
	var s Source
	s.Set(value)
	return s
}

func (s Source) Equals(other Source) bool {
	return s.Value() == other.Value()
}

func (s Source) Verify() error {
	if len(s.Value()) > 100 || len(s.Value()) == 0 {
		return ErrSource
	}
	return nil
}

var (
	ErrSource = fmt.Errorf("source length must be between 1 and 100 characters")
)
