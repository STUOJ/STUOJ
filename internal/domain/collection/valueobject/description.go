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

func (c Description) Equals(other Description) bool {
	return c.Value() == other.Value()
}

func (c Description) Verify() error {
	if len(c.Value()) > 100000 || len(c.Value()) == 0 {
		return ErrContent
	}
	return nil
}

var (
	ErrContent = fmt.Errorf("description length must be between 1 and 100000 characters")
)
