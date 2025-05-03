package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
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
	return i.Value() == other.Value()
}

func (i Input) Verify() error {
	if len(i.Value()) > 100000 || len(i.Value()) == 0 {
		return ErrInput
	}
	return nil
}

var (
	ErrInput = fmt.Errorf("input description length must be between 1 and 100000 characters")
)
