package valueobject

import (
	"STUOJ/internal/domain/shared"
)

// SubmissionID 表示提交记录的唯一标识
// 封装提交ID的领域逻辑和验证规则
type SubmissionId struct {
	shared.Valueobject[int64]
}

// NewSubmissionID 创建新的提交ID值对象
// 参数：value - 必须大于0的整型值
func NewSubmissionId(value int64) SubmissionId {
	var id SubmissionId
	id.Set(value)
	return id
}
