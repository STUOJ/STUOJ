package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Password 只允许通过构造函数创建
type Password struct {
	shared.Valueobject[string]        // 密文，存数据库或从数据库读出
	plaintext                  string // 明文，仅新建/修改时用，其他场景为空
}

// NewPasswordPlaintext 创建新密码（明文），用于注册/修改密码
func NewPasswordPlaintext(plaintext string) Password {
	return Password{plaintext: plaintext}
}

// NewPassword 用于数据库读出（只含密文）
func NewPassword(ciphertext string) Password {
	var p Password
	p.Set(ciphertext)
	return p
}

// Hash 返回加密后的Password（只含密文）
func (p Password) Hash() (Password, error) {
	if p.plaintext == "" {
		return p, nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(p.plaintext), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	var hashed Password
	hashed.Set(string(hash))
	return hashed, nil
}

// Verify 校验明文密码格式
func (p Password) Verify() error {
	if p.plaintext != "" && len(p.plaintext) < 6 || len(p.plaintext) > 20 {
		return errors.New("密码长度必须在6-20个字符之间！")
	}
	return nil
}

// VerifyHash 校验输入密码是否与密文匹配
func (p Password) VerifyHash(plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Value()), []byte(plaintext))
}
