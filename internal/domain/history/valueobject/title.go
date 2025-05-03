package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Title struct {
	model.Valueobject[string]
}

func (t Title) Verify() error {
	if len(t.Value()) == 0 {
		return errors.New("标题不能为空！")
	}
	if len(t.Value()) > 255 {
		return errors.New("标题长度不能超过255个字符！")
	}
	return nil
}

func NewTitle(title string) Title {
	var t Title
	t.Set(title)
	return t
}
