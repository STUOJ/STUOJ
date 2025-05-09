package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Description struct {
	shared.Valueobject[string]
}

func (d Description) Verify() error {
	if utf8.RuneCountInString(d.Value()) > 65535 {
		return errors.New("团队简介长度不能超过65535个字符！")
	}
	return nil
}

func NewDescription(description string) Description {
	var d Description
	d.Set(description)
	return d
}
