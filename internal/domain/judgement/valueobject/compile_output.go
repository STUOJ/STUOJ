package valueobject

import (
	"STUOJ/internal/domain/shared"
)

// CompileOutput 封装判题记录的编译输出内容
// 包含输出内容验证逻辑
type CompileOutput struct {
	shared.Valueobject[string]
}

// NewCompileOutput 创建编译输出值对象
func NewCompileOutput(content string) CompileOutput {
	var co CompileOutput
	co.Set(content)
	return co
}
