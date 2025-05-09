package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
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
	if utf8.RuneCountInString(t.Value()) > 50 || len(t.Value()) == 0 {
		return errors.New("标题应该在50个字符以内，且不能为空")
	}
	return nil
}
