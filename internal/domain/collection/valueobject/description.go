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

func (c Description) String() string {
	return c.value
}

func (c Description) Equals(other Description) bool {
	return c.value == other.value
}

func (c Description) Verify() error {
	if len(c.value) > 100000 || len(c.value) == 0 {
		return ErrContent
	}
	return nil
}

var (
	ErrContent = fmt.Errorf("description length must be between 1 and 100000 characters")
)
