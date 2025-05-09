package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Output struct {
	shared.Valueobject[string]
}

func NewOutput(value string) Output {
	var o Output
	o.Set(value)
	return o
}

func (o Output) Equals(other Output) bool {
	if o.Exist() && other.Exist() {
		return o.Value() == other.Value()
	}
	return false
}

func (o Output) Verify() error {
	if utf8.RuneCountInString(o.Value()) > 100000 {
		return errors.New("output length exceeds the limit")
	}
	return nil
}
