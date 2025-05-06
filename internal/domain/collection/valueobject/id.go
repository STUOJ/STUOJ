package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Id struct {
	model.Valueobject[int64]
}

func NewId(value int64) Id {
	var i Id
	i.Set(value)
	return i
}

func (i Id) Verify() error {
	if i.Value() <= 0 {
		return errors.ErrId
	}
	return nil
}
