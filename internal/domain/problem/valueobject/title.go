package valueobject

import (
	"STUOJ/internal/domain/shared"
	"fmt"
	"unicode/utf8"
)

type Title struct {
	shared.Valueobject[string]
}

func NewTitle(value string) Title {
	var t Title
	t.Set(value)
	return t
}

func (t Title) Equals(other Title) bool {
	if t.Exist() && other.Exist() {
		return t.Value() == other.Value()
	}
	return false
}

func (t Title) Verify() error {
	if utf8.RuneCountInString(t.Value()) > 255 || len(t.Value()) == 0 {
		return ErrTitle
	}
	return nil
}

var (
	ErrTitle = fmt.Errorf("title length must be between 1 and 255 characters")
)
