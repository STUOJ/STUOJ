package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Name struct {
	shared.Valueobject[string]
}

func (n Name) Verify() error {
	if len(n.Value()) == 0 {
		return errors.New("语言名称不能为空！")
	}
	if utf8.RuneCountInString(n.Value()) > 255 {
		return errors.New("语言名称长度不能超过255个字符！")
	}
	return nil
}

func NewName(name string) Name {
	var n Name
	n.Set(name)
	return n
}
