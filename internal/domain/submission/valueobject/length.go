package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type Length struct {
	shared.Valueobject[int64]
}

func NewLength(value int64) Length {
	var s Length
	s.Set(value)
	return s
}

func (s Length) Verify() error {
	if s.Value() <= 0 {
		return errors.New("长度不合法")
	}
	return nil
}
