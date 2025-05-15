package valueobject

import (
	"STUOJ/internal/domain/shared"
	"fmt"
)

type Source struct {
	shared.Valueobject[string]
}

func NewSource(value string) Source {
	var s Source
	s.Set(value)
	return s
}

func (s Source) Equals(other Source) bool {
	if s.Exist() && other.Exist() {
		return s.Value() == other.Value()
	}
	return false
}

func (s Source) Verify() error {
	if len(s.Value()) > 100 {
		return ErrSource
	}
	return nil
}

var (
	ErrSource = fmt.Errorf("来源过长")
)
