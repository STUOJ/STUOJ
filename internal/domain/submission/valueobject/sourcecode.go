package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type SourceCode struct {
	shared.Valueobject[string]
}

func (s SourceCode) Verify() error {
	if len(s.Value()) == 0 {
		return errors.New("源代码不能为空！")
	}
	return nil
}

func NewSourceCode(sourceCode string) SourceCode {
	var s SourceCode
	s.Set(sourceCode)
	return s
}
