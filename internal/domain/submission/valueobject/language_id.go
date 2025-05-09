package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type LanguageId struct {
	shared.Valueobject[int64]
}

func NewLanguageId(value int64) LanguageId {
	var s LanguageId
	s.Set(value)
	return s
}

func (s LanguageId) Verify() error {
	if s.Value() <= 0 {
		return errors.New("语言Id不合法")
	}
	return nil
}
