package valueobject

import (
	"errors"
	"strings"
)

type Avatar string

func (a Avatar) Verify() error {
	if len(a) == 0 {
		return nil
	}
	if !strings.HasPrefix(string(a), "http://") && !strings.HasPrefix(string(a), "https://") {
		return errors.New("头像URL必须以http://或https://开头！")
	}
	return nil
}

func (a Avatar) String() string {
	return string(a)
}

func NewAvatar(avatar string) Avatar {
	return Avatar(avatar)
}
