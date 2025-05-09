package valueobject

import (
	"STUOJ/internal/model"
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
