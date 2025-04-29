package valueobject

import (
	"STUOJ/pkg/errors"
	"strings"
)

type Content struct {
	value string
}

func (c Content) Verify() error {
	if strings.TrimSpace(c.value) == "" {
		return errors.ErrValidation.WithMessage("评论内容不能为空")
	}
	if len(c.value) > 10000 {
		return errors.ErrValidation.WithMessage("评论内容不能超过10000个字符")
	}
	return nil
}

func (c Content) String() string {
	return c.value
}

func NewContent(value string) Content {
	return Content{value: value}
}
