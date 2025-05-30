package valueobject

import (
	"STUOJ/internal/domain/shared"
)

// TestcaseID 表示测试用例的唯一标识
// 封装测试用例ID的验证逻辑和领域行为
type TestcaseId struct {
	shared.Valueobject[int64]
}

func NewTestcaseId(value int64) TestcaseId {
	var id TestcaseId
	id.Set(value)
	return id
}
