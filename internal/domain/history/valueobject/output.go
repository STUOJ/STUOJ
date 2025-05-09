package valueobject

import (
	"STUOJ/internal/domain/shared"
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
