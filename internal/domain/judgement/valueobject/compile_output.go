package valueobject

import (
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

// CompileOutput 封装判题记录的编译输出内容
// 包含输出内容验证逻辑
type CompileOutput struct {
	model.Valueobject[string]
}

// NewCompileOutput 创建编译输出值对象
func NewCompileOutput(content string) CompileOutput {
	var co CompileOutput
	co.Set(content)
	return co
}

// Verify 验证编译输出内容有效性
func (c CompileOutput) Verify() error {
	if len(c.Value()) > 65535 {
		return errors.ErrValidation.WithMessage("编译输出内容超过最大限制")
	}
	return nil
}
