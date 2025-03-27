package valueobject

import (
	"errors"
)

type Name string

func (n Name) Verify() error {
	if len(n) == 0 {
		return errors.New("队名不能为空！")
	}
	if len(n) > 255 {
		return errors.New("队名长度不能超过255个字符！")
	}
	return nil
}

func (n Name) String() string {
	return string(n)
}

func NewName(name string) Name {
	return Name(name)
}
