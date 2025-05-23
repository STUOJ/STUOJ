package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"strings"
	"unicode/utf8"
)

type Content struct {
	shared.Valueobject[string]
}

func (c Content) Verify() error {
	if strings.TrimSpace(c.Value()) == "" {
		return errors.New("评论内容不能为空")
	}
	if utf8.RuneCountInString(c.Value()) > 10000 {
		return errors.New("评论内容不能超过10000个字符")
	}
	return nil
}

func NewContent(value string) Content {
	var c Content
	c.Set(value)
	return c
}
