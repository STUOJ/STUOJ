package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
)

type MapId struct {
	shared.Valueobject[uint32]
}

func NewMapId(value uint32) MapId {
	var id MapId
	id.Set(value)
	return id
}

func (i MapId) Verify() error {
	if i.Value() <= 0 {
		return errors.New("map id 不能为负数")
	}
	return nil
}
