package valueobject

import (
	"fmt"
)

type Source struct {
	value string
}

func NewSource(value string) Source {
	return Source{value: value}
}

func (s Source) String() string {
	return s.value
}

func (s Source) Equals(other Source) bool {
	return s.value == other.value
}

func (s Source) Verify() error {
	if len(s.value) > 100 || len(s.value) == 0 {
		return ErrSource
	}
	return nil
}

var (
	ErrSource = fmt.Errorf("source length must be between 1 and 100 characters")
)
