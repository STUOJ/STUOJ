package valueobject

import (
	"STUOJ/internal/model"
)

type Time struct {
	model.Valueobject[int64]
}

func NewTime(value int64) Time {
	var t Time
	t.Set(value)
	return t
}
