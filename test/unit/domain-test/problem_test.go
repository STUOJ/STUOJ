package domain_test

import (
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/problem/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"testing"
	"time"
)

// 生成随机题目标题
func randomProblemTitle() string {
	return "题目" + time.Now().Format("150405.000")
}
func randomProblemSource() string {
	return "来源" + time.Now().Format("150405.000")
}
func randomProblemDesc() string {
	return "描述" + time.Now().Format("150405.000")
}
func randomProblemInput() string {
	return "输入" + time.Now().Format("150405.000")
}
func randomProblemOutput() string {
	return "输出" + time.Now().Format("150405.000")
}

// 测试题目创建成功
func TestProblemCreate_Success(t *testing.T) {
	p := problem.NewProblem(
		problem.WithTitle(randomProblemTitle()),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	id, err := p.Create()
	if err != nil {
		t.Fatalf("创建题目失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建题目返回的ID无效")
	}
}

// 测试题目标题为空时创建失败
func TestProblemCreate_EmptyTitle(t *testing.T) {
	p := problem.NewProblem(
		problem.WithTitle(""),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	_, err := p.Create()
	if err == nil {
		t.Fatalf("题目标题为空时应创建失败")
	}
}

// 测试题目更新成功
func TestProblemUpdate_Success(t *testing.T) {
	p := problem.NewProblem(
		problem.WithTitle(randomProblemTitle()),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	id, err := p.Create()
	if err != nil {
		t.Fatalf("创建题目失败: %v", err)
	}
	p.Id.Set(id) // 使用Valueobject设置ID
	p.Description = valueobject.NewDescription("更新后的描述")
	err = p.Update()
	if err != nil {
		t.Fatalf("更新题目失败: %v", err)
	}
}

// 测试更新不存在的题目
func TestProblemUpdate_NotFound(t *testing.T) {
	p := problem.NewProblem(
		problem.WithId(99999999),
		problem.WithTitle(randomProblemTitle()),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	err := p.Update()
	if err == nil {
		t.Fatalf("更新不存在的题目应失败")
	}
}

// 测试题目删除成功
func TestProblemDelete_Success(t *testing.T) {
	p := problem.NewProblem(
		problem.WithTitle(randomProblemTitle()),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	id, err := p.Create()
	if err != nil {
		t.Fatalf("创建题目失败: %v", err)
	}
	p.Id.Set(id)
	err = p.Delete()
	if err != nil {
		t.Fatalf("删除题目失败: %v", err)
	}
}

// 测试删除不存在的题目
func TestProblemDelete_NotFound(t *testing.T) {
	p := problem.NewProblem(
		problem.WithId(99999999),
		problem.WithTitle(randomProblemTitle()),
		problem.WithSource(randomProblemSource()),
		problem.WithDifficulty(entity.DifficultyE),
		problem.WithTimeLimit(1.0),
		problem.WithMemoryLimit(128),
		problem.WithDescription(randomProblemDesc()),
		problem.WithInput(randomProblemInput()),
		problem.WithOutput(randomProblemOutput()),
		problem.WithSampleInput(randomProblemInput()),
		problem.WithSampleOutput(randomProblemOutput()),
		problem.WithHint(randomProblemDesc()),
		problem.WithStatus(entity.ProblemPublic),
	)
	err := p.Delete()
	if err == nil {
		t.Fatalf("删除不存在的题目应失败")
	}
}
