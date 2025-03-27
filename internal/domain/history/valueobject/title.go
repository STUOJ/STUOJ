package valueobject

import (
	"errors"
)

type Title string

func (t Title) Verify() error {
	if len(t) == 0 {
		return errors.New("标题不能为空！")
	}
	if len(t) > 255 {
		return errors.New("标题长度不能超过255个字符！")
	}
	return nil
}

func (t Title) String() string {
	return string(t)
}

func NewTitle(title string) Title {
	return Title(title)
}