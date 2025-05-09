package valueobject

import (
	"STUOJ/internal/model"
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
