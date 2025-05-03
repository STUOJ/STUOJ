package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
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

func (d Description) Verify() error {
	if len(d.Value()) > 100000 || len(d.Value()) == 0 {
		return ErrDescription
	}
	return nil
}

var (
	ErrDescription = fmt.Errorf("description length must be between 1 and 100000 characters")
)
