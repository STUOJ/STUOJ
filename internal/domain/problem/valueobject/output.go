package valueobject

import (
	"fmt"
)

type Output struct {
	value string
}

func NewOutput(value string) Output {
	return Output{value: value}
}

func (o Output) String() string {
	return o.value
}

func (o Output) Equals(other Output) bool {
	return o.value == other.value
}

func (o Output) Verify() error {
	if len(o.value) > 100000 || len(o.value) == 0 {
		return ErrOutput
	}
	return nil
}

var (
	ErrOutput = fmt.Errorf("output description length must be between 1 and 100000 characters")
)
