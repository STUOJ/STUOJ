package domain_test

import (
	"STUOJ/internal/domain/solution"
	"STUOJ/internal/domain/solution/valueobject"
	"math/rand"
	"testing"
)

// 生成随机代码
func randomSourceCode() string {
	return "print(\"Hello, World " + string(rune(rand.Intn(26)+'A')) + "\")"
}

// 测试提交创建成功
func TestSolutionCreate_Success(t *testing.T) {
	s := solution.NewSolution(
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(randomSourceCode()),
	)
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
	s := solution.NewSolution(
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(""),
	)
	_, err := s.Create()
	if err == nil {
		t.Fatalf("代码为空时应创建失败")
	}
}

// 测试提交更新成功
func TestSolutionUpdate_Success(t *testing.T) {
	s := solution.NewSolution(
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(randomSourceCode()),
	)
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
	s := solution.NewSolution(
		solution.WithId(99999999),
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(randomSourceCode()),
	)
	err := s.Update()
	if err == nil {
		t.Fatalf("更新不存在的提交应失败")
	}
}

// 测试提交删除成功
func TestSolutionDelete_Success(t *testing.T) {
	s := solution.NewSolution(
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(randomSourceCode()),
	)
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
	s := solution.NewSolution(
		solution.WithId(99999999),
		solution.WithLanguageId(1),
		solution.WithProblemId(1),
		solution.WithSourceCode(randomSourceCode()),
	)
	err := s.Delete()
	if err == nil {
		t.Fatalf("删除不存在的提交应失败")
	}
}
