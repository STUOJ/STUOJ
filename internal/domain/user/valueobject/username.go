package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/utils"
	"errors"
	"strings"
	"unicode/utf8"
)

type Username struct {
	model.Valueobject[string]
}

func (u Username) Verify() error {
	val := u.Value()
	if utf8.RuneCountInString(val) < 3 || utf8.RuneCountInString(val) > 12 {
		return errors.New("用户名长度必须在3-12个字符之间！")
	}
	if strings.ContainsAny(val, " \t\n\r") {
		return errors.New("用户名不能包含空白字符！")
	}
	return nil
}

func NewUsername(un string) Username {
	un = utils.Senitize(un)
	var u Username
	u.Set(un)
	return u
}
