package valueobject

import (
	"fmt"
)

type Content struct {
	value string
}

func NewContent(value string) Content {
	return Content{value: value}
}

func (c Content) String() string {
	return c.value
}

func (c Content) Equals(other Content) bool {
	return c.value == other.value
}

func (c Content) Verify() error {
	if len(c.value) > 100000 || len(c.value) == 0 {
		return ErrContent
	}
	return nil
}

var (
	ErrContent = fmt.Errorf("content length must be between 1 and 100000 characters")
)
