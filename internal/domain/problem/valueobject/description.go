package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"unicode/utf8"
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
	if d.Exist() && other.Exist() {
		return d.Value() == other.Value()
	}
	return false
}

func (d Description) Verify() error {
	if utf8.RuneCountInString(d.Value()) > 100000 {
		return errors.New("描述长度错误")
	}
	return nil
}
