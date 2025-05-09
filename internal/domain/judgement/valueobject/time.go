package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Time struct {
	shared.Valueobject[float64]
}

func NewTime(value float64) Time {
	var t Time
	t.Set(value)
	return t
}
