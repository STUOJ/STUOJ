package valueobject

import (
	"fmt"
)

type Title struct {
	value string
}

func NewTitle(value string) Title {
	return Title{value: value}
}

func (t Title) String() string {
	return t.value
}

func (t Title) Equals(other Title) bool {
	return t.value == other.value
}

func (t Title) Verify() error {
	if len(t.value) > 255 || len(t.value) == 0 {
		return ErrTitle
	}
	return nil
}

var (
	ErrTitle = fmt.Errorf("title length must be between 1 and 255 characters")
)
