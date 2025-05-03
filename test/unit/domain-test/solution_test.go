package domain_test

import (
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/domain/solution/valueobject"
	"math/rand"
	"testing"
)

// 生成随机代码
func randomSolutionSourceCode() valueobject.SourceCode {
	return valueobject.NewSourceCode("print(\"Hello, World " + string(rune(rand.Intn(26)+'A')) + "\")")
}

// 测试提交创建成功
func TestSolutionCreate_Success(t *testing.T) {
	s := &solution.Solution{
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: randomSolutionSourceCode(),
	}
	id, err := s.Create()
	if err != nil {
		t.Fatalf("创建提交失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建提交返回的ID无效")
	}
}

// 测试代码为空时创建失败
func TestSolutionCreate_EmptySourceCode(t *testing.T) {
	s := &solution.Solution{
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: valueobject.NewSourceCode(""),
	}
	_, err := s.Create()
	if err == nil {
		t.Fatalf("代码为空时应创建失败")
	}
}

// 测试提交更新成功
func TestSolutionUpdate_Success(t *testing.T) {
	s := &solution.Solution{
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: randomSolutionSourceCode(),
	}
	id, err := s.Create()
	if err != nil {
		t.Fatalf("创建提交失败: %v", err)
	}
	s.Id = id
	s.SourceCode = valueobject.NewSourceCode("print(\"Updated\")")
	err = s.Update()
	if err != nil {
		t.Fatalf("更新提交失败: %v", err)
	}
}

// 测试更新不存在的提交
func TestSolutionUpdate_NotFound(t *testing.T) {
	s := &solution.Solution{
		Id:         99999999,
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: randomSolutionSourceCode(),
	}
	err := s.Update()
	if err == nil {
		t.Fatalf("更新不存在的提交应失败")
	}
}

// 测试提交删除成功
func TestSolutionDelete_Success(t *testing.T) {
	s := &solution.Solution{
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: randomSolutionSourceCode(),
	}
	id, err := s.Create()
	if err != nil {
		t.Fatalf("创建提交失败: %v", err)
	}
	s.Id = id
	err = s.Delete()
	if err != nil {
		t.Fatalf("删除提交失败: %v", err)
	}
}

// 测试删除不存在的提交
func TestSolutionDelete_NotFound(t *testing.T) {
	s := &solution.Solution{
		Id:         99999999,
		LanguageId: 1,
		ProblemId:  1,
		SourceCode: randomSolutionSourceCode(),
	}
	err := s.Delete()
	if err == nil {
		t.Fatalf("删除不存在的提交应失败")
	}
}
