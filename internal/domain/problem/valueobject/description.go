package valueobject

import (
	"fmt"
)

type Description struct {
	value string
}

func NewDescription(value string) Description {
	return Description{value: value}
}

func (d Description) String() string {
	return d.value
}

func (d Description) Equals(other Description) bool {
	return d.value == other.value
}

func (d Description) Verify() error {
	if len(d.value) > 100000 || len(d.value) == 0 {
		return ErrDescription
	}
	return nil
}

var (
	ErrDescription = fmt.Errorf("description length must be between 1 and 100000 characters")
)
