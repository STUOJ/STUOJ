package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"strings"
)

type SourceCode struct {
	model.Valueobject[string]
}

func (s SourceCode) Verify() error {
	if strings.TrimSpace(s.Value()) == "" {
		return errors.ErrValidation.WithMessage("源代码不能为空")
	}
	if len(s.Value()) > 100000 {
		return errors.ErrValidation.WithMessage("源代码不能超过100000个字符")
	}
	return nil
}

func NewSourceCode(value string) SourceCode {
	var s SourceCode
	s.Set(value)
	return s
}
