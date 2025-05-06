package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"strings"
)

type Content struct {
	model.Valueobject[string]
}

func (c Content) Verify() error {
	if strings.TrimSpace(c.Value()) == "" {
		return errors.New("评论内容不能为空")
	}
	if len(c.Value()) > 10000 {
		return errors.New("评论内容不能超过10000个字符")
	}
	return nil
}

func NewContent(value string) Content {
	var c Content
	c.Set(value)
	return c
}
