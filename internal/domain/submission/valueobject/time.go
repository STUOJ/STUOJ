package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Time struct {
	shared.Valueobject[int64]
}

func NewTime(value int64) Time {
	var t Time
	t.Set(value)
	return t
}
