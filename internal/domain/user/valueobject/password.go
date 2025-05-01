package valueobject

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	plaintext  string
	ciphertext string
}

func (p Password) Verify() error {
	if len(p.plaintext) < 6 || len(p.plaintext) > 20 {
		return errors.New("密码长度必须在6-20个字符之间！")
	}
	return nil
}

func (p Password) String() string {
	return p.plaintext
}

func hash(plaintext string) (string, error) {
	ciphertext, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

// Hash 密码加密
func (p Password) Hash() (Password, error) {
	var err error

	if p.plaintext == "" {
		return Password{}, errors.New("密码不能为空")
	}
	p.ciphertext, err = hash(p.plaintext)
	if err != nil {
		return Password{}, err
	}

	return p, nil
}

// VerifyHash 验证密码
func (p Password) VerifyHash(pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.ciphertext), []byte(pw))
}

func NewPassword(plaintext string) Password {
	return Password{
		plaintext:  plaintext,
		ciphertext: "",
	}
}
