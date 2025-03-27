package valueobject

import (
	"STUOJ/internal/errors"
)

type TestOutput struct {
	value string
}

func (t TestOutput) Verify() error {
	if len(t.value) > 100000 {
		return errors.ErrValidation.WithMessage("测试输出不能超过100000个字符")
	}
	return nil
}

func (t TestOutput) String() string {
	return t.value
}

func NewTestOutput(value string) TestOutput {
	return TestOutput{value: value}
}
