package valueobject

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p Password) Verify() error {
	if len(p) < 6 || len(p) > 20 {
		return errors.New("密码长度必须在6-20个字符之间！")
	}
	return nil
}

func (p Password) String() string {
	return string(p)
}

// Hash 密码加密
func (p Password) Hash() (Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return Password(hashedPassword), nil
}

// VerifyHash 验证密码
func (p Password) VerifyHash(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(pw))
}

func NewPassword(pw string) Password {
	return Password(pw)
}
