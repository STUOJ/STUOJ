package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"unicode/utf8"
)

type Input struct {
	model.Valueobject[string]
}

func NewInput(value string) Input {
	var i Input
	i.Set(value)
	return i
}

func (i Input) Equals(other Input) bool {
	if i.Exist() && other.Exist() {
		return i.Value() == other.Value()
	}
	return false
}

func (i Input) Verify() error {
	if utf8.RuneCountInString(i.Value()) > 100000 {
		return errors.New("input too long")
	}
	return nil
}
