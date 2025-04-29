package valueobject

import (
	"STUOJ/pkg/utils"
	"errors"
)

type Signature string

func (s Signature) Verify() error {
	if len(s) > 200 {
		return errors.New("个性签名长度不能超过200个字符！")
	}
	return nil
}

func (s Signature) String() string {
	return string(s)
}

func NewSignature(s string) Signature {
	// 预处理
	s = utils.Senitize(s)

	return Signature(s)
}
