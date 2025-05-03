package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Name struct {
	model.Valueobject[string]
}

func (n Name) Verify() error {
	if len(n.Value()) == 0 {
		return errors.New("队名不能为空！")
	}
	if len(n.Value()) > 255 {
		return errors.New("队名长度不能超过255个字符！")
	}
	return nil
}

func NewName(name string) Name {
	var n Name
	n.Set(name)
	return n
}
