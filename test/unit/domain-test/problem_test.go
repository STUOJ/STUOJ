package domain_test

import (
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/domain/problem/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"testing"
	"time"
)

// 生成随机题目标题
func randomProblemTitle() valueobject.Title {
	return valueobject.NewTitle("题目" + time.Now().Format("150405.000"))
}
func randomProblemSource() valueobject.Source {
	return valueobject.NewSource("来源" + time.Now().Format("150405.000"))
}
func randomProblemDesc() valueobject.Description {
	return valueobject.NewDescription("描述" + time.Now().Format("150405.000"))
}
func randomProblemInput() valueobject.Input {
	return valueobject.NewInput("输入" + time.Now().Format("150405.000"))
}
func randomProblemOutput() valueobject.Output {
	return valueobject.NewOutput("输出" + time.Now().Format("150405.000"))
}

// 测试题目创建成功
func TestProblemCreate_Success(t *testing.T) {
	p := &problem.Problem{
		Title:        randomProblemTitle(),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
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
	p := &problem.Problem{
		Title:        valueobject.NewTitle(""),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
	_, err := p.Create()
	if err == nil {
		t.Fatalf("题目标题为空时应创建失败")
	}
}

// 测试题目更新成功
func TestProblemUpdate_Success(t *testing.T) {
	p := &problem.Problem{
		Title:        randomProblemTitle(),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
	id, err := p.Create()
	if err != nil {
		t.Fatalf("创建题目失败: %v", err)
	}
	p.Id = id
	p.Description = valueobject.NewDescription("更新后的描述")
	err = p.Update()
	if err != nil {
		t.Fatalf("更新题目失败: %v", err)
	}
}

// 测试更新不存在的题目
func TestProblemUpdate_NotFound(t *testing.T) {
	p := &problem.Problem{
		Id:           99999999,
		Title:        randomProblemTitle(),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
	err := p.Update()
	if err == nil {
		t.Fatalf("更新不存在的题目应失败")
	}
}

// 测试题目删除成功
func TestProblemDelete_Success(t *testing.T) {
	p := &problem.Problem{
		Title:        randomProblemTitle(),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
	id, err := p.Create()
	if err != nil {
		t.Fatalf("创建题目失败: %v", err)
	}
	p.Id = id
	err = p.Delete()
	if err != nil {
		t.Fatalf("删除题目失败: %v", err)
	}
}

// 测试删除不存在的题目
func TestProblemDelete_NotFound(t *testing.T) {
	p := &problem.Problem{
		Id:           99999999,
		Title:        randomProblemTitle(),
		Source:       randomProblemSource(),
		Difficulty:   entity.DifficultyE,
		TimeLimit:    1.0,
		MemoryLimit:  128,
		Description:  randomProblemDesc(),
		Input:        randomProblemInput(),
		Output:       randomProblemOutput(),
		SampleInput:  randomProblemInput(),
		SampleOutput: randomProblemOutput(),
		Hint:         randomProblemDesc(),
		Status:       entity.ProblemPublic,
	}
	err := p.Delete()
	if err == nil {
		t.Fatalf("删除不存在的题目应失败")
	}
}
