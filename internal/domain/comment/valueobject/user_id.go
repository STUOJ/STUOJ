package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// UserID 表示评论关联的用户ID值对象
type UserID struct {
	model.Valueobject[int64]
}

// Verify 验证用户ID是否有效
func (u UserID) Verify() error {
	if u.Value() <= 0 {
		return errors.ErrUserId
	}
	return nil
}

// NewUserID 创建一个新的用户ID值对象
func NewUserID(id int64) UserID {
	var u UserID
	u.Set(id)
	return u
}
