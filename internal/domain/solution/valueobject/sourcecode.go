package valueobject

import (
	"STUOJ/pkg/errors"
	"strings"
)

type SourceCode struct {
	value string
}

func (s SourceCode) Verify() error {
	if strings.TrimSpace(s.value) == "" {
		return errors.ErrValidation.WithMessage("源代码不能为空")
	}
	if len(s.value) > 100000 {
		return errors.ErrValidation.WithMessage("源代码不能超过100000个字符")
	}
	return nil
}

func (s SourceCode) String() string {
	return s.value
}

func NewSourceCode(value string) SourceCode {
	return SourceCode{value: value}
}
