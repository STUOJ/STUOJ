package valueobject

import (
	"STUOJ/internal/model"
	"errors"
)

// TestcaseID 表示测试用例的唯一标识
// 封装测试用例ID的验证逻辑和领域行为
type TestcaseId struct {
	model.Valueobject[int64]
}

func NewTestcaseId(value int64) TestcaseId {
	var id TestcaseId
	id.Set(value)
	return id
}

// Verify 验证测试用例ID是否合法
func (t TestcaseId) Verify() error {
	if t.Value() <= 0 {
		return errors.New("testcase_id must be greater than 0")
	}
	return nil
}
