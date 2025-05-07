package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"unicode/utf8"
)

type Name struct {
	model.Valueobject[string]
}

func (n Name) Verify() error {
	if len(n.Value()) == 0 {
		return errors.New("队名不能为空！")
	}
	if utf8.RuneCountInString(n.Value()) > 12 {
		return errors.New("队名长度不能超过12个字符！")
	}
	return nil
}

func NewName(name string) Name {
	var n Name
	n.Set(name)
	return n
}
