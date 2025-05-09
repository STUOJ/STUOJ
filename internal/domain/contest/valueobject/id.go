package valueobject

import (
	"STUOJ/internal/model"
	"fmt"
)

type Id struct {
	model.Valueobject[int64]
}

func NewId(value int64) Id {
	var id Id
	id.Set(value)
	return id
}

func (id Id) Verify() error {
	if id.Value() <= 0 {
		return ErrId
	}
	return nil
}

var ErrId = fmt.Errorf("id is invalid")
