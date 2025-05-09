package valueobject

import (
	"STUOJ/internal/domain/shared"
)

type Memory struct {
	shared.Valueobject[int64]
}

func NewMemory(value int64) Memory {
	var mem Memory
	mem.Set(value)
	return mem
}
