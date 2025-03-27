package valueobject

import (
	"errors"
	"strings"
)

type Username string

func (u Username) Verify() error {
	if len(u) < 3 || len(u) > 20 {
		return errors.New("用户名长度必须在3-20个字符之间！")
	}
	if strings.ContainsAny(string(u), " \t\n\r") {
		return errors.New("用户名不能包含空白字符！")
	}
	return nil
}

func (u Username) String() string {
	return string(u)
}

func NewUsername(username string) Username {
	return Username(username)
}
