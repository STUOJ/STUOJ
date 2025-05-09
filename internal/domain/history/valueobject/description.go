package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Description struct {
	shared.Valueobject[string]
}

func NewDescription(value string) Description {
	var d Description
	d.Set(value)
	return d
}

func (d Description) Equals(other Description) bool {
	return d.Value() == other.Value()
}
