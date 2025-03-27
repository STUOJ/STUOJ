package valueobject

import (
	"errors"
)

type Title string

func (t Title) Verify() error {
	if len(t) < 3 || len(t) > 50 {
		return errors.New("比赛标题长度必须在3-50个字符之间！")
	}
	return nil
}

func (t Title) String() string {
	return string(t)
}

func NewTitle(title string) Title {
	return Title(title)
}
