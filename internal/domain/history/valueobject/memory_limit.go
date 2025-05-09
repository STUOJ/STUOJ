package valueobject

import (
	"STUOJ/internal/domain/shared"
)

// MemoryLimit 表示历史记录中的内存限制值对象
type MemoryLimit struct {
	shared.Valueobject[int64]
}

// NewMemoryLimit 创建一个新的MemoryLimit值对象
func NewMemoryLimit(value int64) MemoryLimit {
	var m MemoryLimit
	m.Set(value)
	return m
}
