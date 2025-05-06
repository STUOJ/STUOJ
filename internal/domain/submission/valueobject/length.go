package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Length struct {
	model.Valueobject[int64]
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
