package valueobject

import (
	"errors"
	"regexp"
)

type Email string

func (e Email) Verify() error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(string(e)) {
		return errors.New("邮箱格式不正确！")
	}
	return nil
}

func (e Email) String() string {
	return string(e)
}

func NewEmail(email string) Email {
	return Email(email)
}
