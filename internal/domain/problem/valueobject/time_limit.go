package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type TimeLimit struct {
	model.Valueobject[float64]
}

func NewTimeLimit(value float64) TimeLimit {
	var tl TimeLimit
	tl.Set(value)
	return tl
}

func (tl TimeLimit) Verify() error {
	if tl.Value() <= 0 {
		return errors.New("时间限制错误")
	}
	return nil
}
