package valueobject

import (
	"STUOJ/internal/errors"
)

type TestInput struct {
	value string
}

func (t TestInput) Verify() error {
	if len(t.value) > 100000 {
		return errors.ErrValidation.WithMessage("测试输入不能超过100000个字符")
	}
	return nil
}

func (t TestInput) String() string {
	return t.value
}

func NewTestInput(value string) TestInput {
	return TestInput{value: value}
}
