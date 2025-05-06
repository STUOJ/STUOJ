package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Id struct {
	model.Valueobject[int64]
}

func NewId(value int64) Id {
	var id Id
	id.Set(value)
	return id
}

func (i Id) Verify() error {
	if i.Value() <= 0 {
		return errors.ErrId
	}
	return nil
}
