package valueobject

import (
	"STUOJ/internal/model"
)

// ProblemId 表示历史记录关联的题目ID值对象
type ProblemId struct {
	model.Valueobject[int64]
}

// NewProblemId 创建一个新的ProblemId值对象
func NewProblemId(value int64) ProblemId {
	var p ProblemId
	p.Set(value)
	return p
}
