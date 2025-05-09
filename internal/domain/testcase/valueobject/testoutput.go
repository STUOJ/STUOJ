package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/pkg/errors"
)

type TestOutput struct {
	shared.Valueobject[string]
}

func (t TestOutput) Verify() error {
	if len(t.Value()) > 100000 {
		return errors.ErrValidation.WithMessage("测试输出不能超过100000个字符")
	}
	return nil
}

func NewTestOutput(value string) TestOutput {
	var t TestOutput
	t.Set(value)
	return t
}
