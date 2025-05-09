package valueobject

import (
	"STUOJ/internal/model"
)

type Description struct {
	model.Valueobject[string]
}

func NewDescription(value string) Description {
	var d Description
	d.Set(value)
	return d
}

func (d Description) Equals(other Description) bool {
	return d.Value() == other.Value()
}
