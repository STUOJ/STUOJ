package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

// SubmissionID 表示提交记录的唯一标识
// 封装提交ID的领域逻辑和验证规则
type SubmissionId struct {
	model.Valueobject[int64]
}

// NewSubmissionID 创建新的提交ID值对象
// 参数：value - 必须大于0的整型值
func NewSubmissionId(value int64) SubmissionId {
	var id SubmissionId
	id.Set(value)
	return id
}

func (id SubmissionId) Verify() error {
	if id.Value() <= 0 {
		return errors.New("invalid submission id")
	}
	return nil
}
