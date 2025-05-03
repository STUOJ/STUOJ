package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
	"strings"
)

type Name struct {
	model.Valueobject[string]
}

func (n Name) Verify() error {
	if strings.TrimSpace(n.Value()) == "" {
		return errors.ErrValidation.WithMessage("标签名不能为空")
	}
	if len(n.Value()) > 50 {
		return errors.ErrValidation.WithMessage("标签名不能超过50个字符")
	}
	return nil
}

func NewName(value string) Name {
	var n Name
	n.Set(value)
	return n
}
