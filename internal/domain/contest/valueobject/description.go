package valueobject

import (
	"errors"
)

type Description string

func (d Description) Verify() error {
	if len(d) < 10 || len(d) > 5000 {
		return errors.New("比赛描述长度必须在10-5000个字符之间！")
	}
	return nil
}

func (d Description) String() string {
	return string(d)
}

func NewDescription(description string) Description {
	return Description(description)
}
