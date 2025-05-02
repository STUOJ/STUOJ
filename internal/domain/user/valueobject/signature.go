package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/utils"
	"errors"
)

type Signature struct {
	model.Valueobject[string]
}

func (s Signature) Verify() error {
	if len(s.Value()) > 200 {
		return errors.New("个性签名长度不能超过200个字符！")
	}
	return nil
}

func NewSignature(str string) Signature {
	str = utils.Senitize(str)
	var s Signature
	s.Set(str)
	return s
}
