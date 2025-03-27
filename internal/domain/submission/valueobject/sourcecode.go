package valueobject

import (
	"errors"
)

type SourceCode string

func (s SourceCode) Verify() error {
	if len(s) == 0 {
		return errors.New("源代码不能为空！")
	}
	return nil
}

func (s SourceCode) String() string {
	return string(s)
}

func NewSourceCode(sourceCode string) SourceCode {
	return SourceCode(sourceCode)
}
