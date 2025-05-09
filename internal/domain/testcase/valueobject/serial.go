package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type Serial struct {
	shared.Valueobject[uint16]
}

func NewSerial(value uint16) Serial {
	var serial Serial
	serial.Set(value)
	return serial
}

func (s Serial) Verify() error {
	if s.Value() <= 0 {
		return errors.New("serial 不能为负数")
	}
	return nil
}
