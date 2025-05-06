package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

// BlogID 表示评论关联的博客ID值对象
type BlogID struct {
	model.Valueobject[int64]
}

// Verify 验证博客ID是否有效
func (b BlogID) Verify() error {
	if b.Value() <= 0 {
		return errors.New("博客ID无效")
	}
	return nil
}

// NewBlogID 创建一个新的博客ID值对象
func NewBlogID(id int64) BlogID {
	var b BlogID
	b.Set(id)
	return b
}
