package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"unicode/utf8"
)

type Description struct {
	model.Valueobject[string]
}

func (d Description) Verify() error {
	if utf8.RuneCountInString(d.Value()) < 10 || len(d.Value()) > 5000 {
		return errors.New("比赛描述长度必须在10-5000个字符之间！")
	}
	return nil
}

func NewDescription(description string) Description {
	var d Description
	d.Set(description)
	return d
}
