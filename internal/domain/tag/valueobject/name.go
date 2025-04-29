package valueobject

import (
	"STUOJ/pkg/errors"
	"strings"
)

type Name struct {
	value string
}

func (n Name) Verify() error {
	if strings.TrimSpace(n.value) == "" {
		return errors.ErrValidation.WithMessage("标签名不能为空")
	}
	if len(n.value) > 50 {
		return errors.ErrValidation.WithMessage("标签名不能超过50个字符")
	}
	return nil
}

func (n Name) String() string {
	return n.value
}

func NewName(value string) Name {
	return Name{value: value}
}
