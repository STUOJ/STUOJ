package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

// Stdout 表示判题标准输出内容
// 封装输出内容的验证逻辑（最大长度限制）
type Stdout struct {
	model.Valueobject[string]
}

func NewStdout(content string) Stdout {
	var so Stdout
	so.Set(content)
	return so
}

func (s Stdout) Verify() error {
	if len(s.Value()) > 65535 {
		return errors.New("stdout 太长")
	}
	return nil
}
