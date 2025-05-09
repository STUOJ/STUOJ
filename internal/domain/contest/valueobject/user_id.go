package valueobject

import (
	"STUOJ/internal/domain/shared"
	"fmt"
)

// UserID 表示比赛关联的用户ID值对象
type UserID struct {
	shared.Valueobject[int64]
}

// Verify 验证用户ID是否有效
func (u UserID) Verify() error {
	if u.Value() <= 0 {
		return fmt.Errorf("用户ID必须大于0")
	}
	return nil
}

// NewUserID 创建一个新的用户ID值对象
func NewUserID(id int64) UserID {
	var u UserID
	u.Set(id)
	return u
}
