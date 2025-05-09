package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type UserId struct {
	model.Valueobject[int64]
}

func NewUserId(value int64) UserId {
	var i UserId
	i.Set(value)
	return i
}

func (i UserId) Verify() error {
	if i.Value() <= 0 {
		return errors.ErrId
	}
	return nil
}
