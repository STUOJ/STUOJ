package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Description struct {
	model.Valueobject[string]
}

func (d Description) Verify() error {
	if len(d.Value()) > 65535 {
		return errors.New("团队简介长度不能超过65535个字符！")
	}
	return nil
}

func NewDescription(description string) Description {
	var d Description
	d.Set(description)
	return d
}
