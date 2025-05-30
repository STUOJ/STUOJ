package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Content struct {
	shared.Valueobject[string]
}

func NewContent(value string) Content {
	var c Content
	c.Set(value)
	return c
}

func (c Content) Verify() error {
	if utf8.RuneCountInString(c.Value()) > 100000 || len(c.Value()) == 0 {
		return errors.New("内容应该在100000个字符以内，且不能为空")
	}
	return nil
}
