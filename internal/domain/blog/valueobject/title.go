package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
)

type Title struct {
	model.Valueobject[string]
}

func NewTitle(value string) Title {
	var t Title
	t.Set(value)
	return t
}

func (t Title) Equals(other Title) bool {
	return t.Value() == other.Value()
}

func (t Title) Verify() error {
	if len(t.Value()) > 50 || len(t.Value()) == 0 {
		return ErrTitle
	}
	return nil
}

var (
	ErrTitle = fmt.Errorf("title length must be between 1 and 50 characters")
)
