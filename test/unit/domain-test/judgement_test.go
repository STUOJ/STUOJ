package domain_test

import (
	"STUOJ/internal/domain/judgement"
	"STUOJ/internal/infrastructure/repository/entity"
	"math/rand"
	"testing"
)

// 生成随机ID
func randomId() int64 {
	return int64(rand.Intn(100000) + 1)
}

// 测试评测记录创建成功
func TestJudgementCreate_Success(t *testing.T) {
	j := &judgement.Judgement{
		SubmissionId:  randomId(),
		TestcaseId:    randomId(),
		Time:          1.23,
		Memory:        128,
		Stdout:        "stdout",
		Stderr:        "",
		CompileOutput: "",
		Message:       "ok",
		Status:        entity.JudgeAC,
	}
	id, err := j.Create()
	if err != nil {
		t.Fatalf("创建评测记录失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建评测记录返回的ID无效")
	}
}

// 测试SubmissionId为空时创建失败
func TestJudgementCreate_EmptySubmissionId(t *testing.T) {
	j := &judgement.Judgement{
		SubmissionId: 0,
		TestcaseId:   randomId(),
		Time:         1.23,
		Memory:       128,
		Status:       entity.JudgeAC,
	}
	_, err := j.Create()
	if err == nil {
		t.Fatalf("SubmissionId为空时应创建失败")
	}
}

// 测试评测记录更新成功
func TestJudgementUpdate_Success(t *testing.T) {
	j := &judgement.Judgement{
		SubmissionId: randomId(),
		TestcaseId:   randomId(),
		Time:         1.23,
		Memory:       128,
		Stdout:       "stdout",
		Status:       entity.JudgeAC,
	}
	id, err := j.Create()
	if err != nil {
		t.Fatalf("创建评测记录失败: %v", err)
	}
	j.Id = id
	j.Message = "updated"
	err = j.Update()
	if err != nil {
		t.Fatalf("更新评测记录失败: %v", err)
	}
}

// 测试更新不存在的评测记录
func TestJudgementUpdate_NotFound(t *testing.T) {
	j := &judgement.Judgement{
		Id:           99999999,
		SubmissionId: randomId(),
		TestcaseId:   randomId(),
		Time:         1.23,
		Memory:       128,
		Status:       entity.JudgeAC,
	}
	err := j.Update()
	if err == nil {
		t.Fatalf("更新不存在的评测记录应失败")
	}
}

// 测试评测记录删除成功
func TestJudgementDelete_Success(t *testing.T) {
	j := &judgement.Judgement{
		SubmissionId: randomId(),
		TestcaseId:   randomId(),
		Time:         1.23,
		Memory:       128,
		Status:       entity.JudgeAC,
	}
	id, err := j.Create()
	if err != nil {
		t.Fatalf("创建评测记录失败: %v", err)
	}
	j.Id = id
	err = j.Delete()
	if err != nil {
		t.Fatalf("删除评测记录失败: %v", err)
	}
}

// 测试删除不存在的评测记录
func TestJudgementDelete_NotFound(t *testing.T) {
	j := &judgement.Judgement{
		Id:           99999999,
		SubmissionId: randomId(),
		TestcaseId:   randomId(),
		Time:         1.23,
		Memory:       128,
		Status:       entity.JudgeAC,
	}
	err := j.Delete()
	if err == nil {
		t.Fatalf("删除不存在的评测记录应失败")
	}
}
