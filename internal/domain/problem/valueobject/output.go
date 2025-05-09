package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"unicode/utf8"
)

type Output struct {
	model.Valueobject[string]
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
