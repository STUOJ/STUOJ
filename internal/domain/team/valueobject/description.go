package valueobject

import (
	"errors"
)

type Description string

func (d Description) Verify() error {
	if len(d) > 65535 {
		return errors.New("团队简介长度不能超过65535个字符！")
	}
	return nil
}

func (d Description) String() string {
	return string(d)
}

func NewDescription(description string) Description {
	return Description(description)
}
