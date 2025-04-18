package valueobject

import (
	"STUOJ/utils"
	"errors"
)

type Signature string

func (s Signature) Verify() error {
	if len(s) > 200 {
		return errors.New("个性签名长度不能超过200个字符！")
	}
	return nil
}

func (s Signature) Sanitize() Signature {
	return Signature(utils.Senitize(string(s)))
}

func (s Signature) String() string {
	return string(s)
}

func NewSignature(signature string) Signature {
	return Signature(signature)
}
