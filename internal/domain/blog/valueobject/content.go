package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
)

type Content struct {
	model.Valueobject[string]
}

func NewContent(value string) Content {
	var c Content
	c.Set(value)
	return c
}

func (c Content) Equals(other Content) bool {
	return c.Value() == other.Value()
}

func (c Content) Verify() error {
	if len(c.Value()) > 100000 || len(c.Value()) == 0 {
		return ErrContent
	}
	return nil
}

var (
	ErrContent = fmt.Errorf("content length must be between 1 and 100000 characters")
)
