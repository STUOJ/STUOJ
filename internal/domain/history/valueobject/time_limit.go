package valueobject

import (
	"STUOJ/internal/model"
)

// TimeLimit 表示历史记录中的时间限制值对象
type TimeLimit struct {
	model.Valueobject[float64]
}

// NewTimeLimit 创建一个新的TimeLimit值对象
func NewTimeLimit(value float64) TimeLimit {
	var t TimeLimit
	t.Set(value)
	return t
}
