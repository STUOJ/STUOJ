package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type MemoryLimit struct {
	model.Valueobject[int64]
}

func NewMemoryLimit(value int64) MemoryLimit {
	var ml MemoryLimit
	ml.Set(value)
	return ml
}

func (ml MemoryLimit) Verify() error {
	if ml.Value() <= 0 {
		return errors.New("内存限制错误")
	}
	return nil
}
