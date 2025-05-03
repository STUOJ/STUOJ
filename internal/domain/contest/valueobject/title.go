package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

type Title struct {
	model.Valueobject[string]
}

func (t Title) Verify() error {
	if len(t.Value()) < 3 || len(t.Value()) > 50 {
		return errors.New("比赛标题长度必须在3-50个字符之间！")
	}
	return nil
}

func NewTitle(title string) Title {
	var t Title
	t.Set(title)
	return t
}
