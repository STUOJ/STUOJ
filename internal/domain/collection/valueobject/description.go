package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Description struct {
	shared.Valueobject[string]
}

func NewDescription(value string) Description {
	var d Description
	d.Set(value)
	return d
}

func (c Description) Equals(other Description) bool {
	if c.Exist() && other.Exist() {
		return c.Value() == other.Value()
	}
	return false
}

func (c Description) Verify() error {
	if utf8.RuneCountInString(c.Value()) > 100000 || len(c.Value()) == 0 {
		return errors.New("内容应该在100000个字符以内，且不能为空")
	}
	return nil
}
