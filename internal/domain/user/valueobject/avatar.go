package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"strings"
)

type Avatar struct {
	model.Valueobject[string]
}

func (a Avatar) Verify() error {
	if !a.Exist() || len(a.Value()) == 0 {
		return nil
	}
	if !strings.HasPrefix(a.Value(), "http://") && !strings.HasPrefix(a.Value(), "https://") {
		return errors.New("头像URL必须以http://或https://开头！")
	}
	return nil
}

func NewAvatar(avatar string) Avatar {
	var a Avatar
	a.Set(avatar)
	return a
}
