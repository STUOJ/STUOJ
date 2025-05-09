package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type TestInput struct {
	model.Valueobject[string]
}

func (t TestInput) Verify() error {
	if len(t.Value()) > 100000 {
		return errors.ErrValidation.WithMessage("测试输入不能超过100000个字符")
	}
	return nil
}

func NewTestInput(value string) TestInput {
	var t TestInput
	t.Set(value)
	return t
}
