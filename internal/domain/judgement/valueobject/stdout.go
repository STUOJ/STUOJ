package valueobject

import (
	"STUOJ/internal/model"
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
