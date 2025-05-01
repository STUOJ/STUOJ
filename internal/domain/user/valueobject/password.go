package valueobject

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Password 只允许通过构造函数创建
type Password struct {
	plaintext  string // 仅新建/修改时用
	ciphertext string // 存数据库或从数据库读出
}

// NewPassword 创建新密码（明文），用于注册/修改密码
func NewPassword(plaintext string) Password {
	return Password{plaintext: plaintext}
}

// NewHashedPassword 用于数据库读出（只含密文）
func NewHashedPassword(ciphertext string) Password {
	return Password{ciphertext: ciphertext}
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
	return Password{ciphertext: string(hash)}, nil
}

// Verify 校验明文密码格式
func (p Password) Verify() error {
	if len(p.plaintext) < 6 || len(p.plaintext) > 20 {
		return errors.New("密码长度必须在6-20个字符之间！")
	}
	return nil
}

// VerifyHash 校验输入密码是否与密文匹配
func (p Password) VerifyHash(plaintext string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.ciphertext), []byte(plaintext))
}

// String 返回密文（用于写入数据库）
func (p Password) String() string {
	return p.ciphertext
}
