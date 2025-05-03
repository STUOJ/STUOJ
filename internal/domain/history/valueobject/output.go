package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
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
	if len(o.Value()) > 100000 || len(o.Value()) == 0 {
		return ErrOutput
	}
	return nil
}

var (
	ErrOutput = fmt.Errorf("output description length must be between 1 and 100000 characters")
)
