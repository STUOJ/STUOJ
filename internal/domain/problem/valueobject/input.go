package valueobject

import (
	"fmt"
)

type Input struct {
	value string
}

func NewInput(value string) Input {
	return Input{value: value}
}

func (i Input) String() string {
	return i.value
}

func (i Input) Equals(other Input) bool {
	return i.value == other.value
}

func (i Input) Verify() error {
	if len(i.value) > 100000 || len(i.value) == 0 {
		return ErrInput
	}
	return nil
}

var (
	ErrInput = fmt.Errorf("input description length must be between 1 and 100000 characters")
)
