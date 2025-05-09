package valueobject

import (
	"STUOJ/internal/domain/shared"
)

// UserId 表示历史记录关联的用户ID值对象
type UserId struct {
	shared.Valueobject[int64]
}

// NewUserId 创建一个新的UserId值对象
func NewUserId(value int64) UserId {
	var u UserId
	u.Set(value)
	return u
}
