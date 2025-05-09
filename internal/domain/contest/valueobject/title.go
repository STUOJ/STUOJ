package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"unicode/utf8"
)

type Title struct {
	shared.Valueobject[string]
}

func (t Title) Verify() error {
	if utf8.RuneCountInString(t.Value()) < 3 || len(t.Value()) > 50 {
		return errors.New("比赛标题长度必须在3-50个字符之间！")
	}
	return nil
}

func NewTitle(title string) Title {
	var t Title
	t.Set(title)
	return t
}
