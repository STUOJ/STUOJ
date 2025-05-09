package valueobject

import (
	"STUOJ/internal/model"
)

type Memory struct {
	model.Valueobject[int64]
}

func NewMemory(value int64) Memory {
	var mem Memory
	mem.Set(value)
	return mem
}
