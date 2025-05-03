package domain_test

import (
	"STUOJ/internal/domain/submission"
	"STUOJ/internal/infrastructure/repository/entity"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机源代码
func randomSourceCode() string {
	return "print(\"Hello, World " + strconv.Itoa(rand.Intn(100000)) + "\")"
}

// 测试提交创建成功
func TestSubmissionCreate_Success(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(20),
		submission.WithLanguageId(1),
		submission.WithSourceCode(randomSourceCode()),
	)
	id, err := s.Create()
	if err != nil {
		t.Fatalf("创建提交失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建提交返回的ID无效")
	}
}

// 测试源代码为空时创建失败
func TestSubmissionCreate_EmptySourceCode(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(0),
		submission.WithLanguageId(1),
		submission.WithSourceCode(""),
	)
	_, err := s.Create()
	if err == nil {
		t.Fatalf("源代码为空时应创建失败")
	}
}

// 测试提交更新成功
func TestSubmissionUpdate_Success(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(20),
		submission.WithLanguageId(1),
		submission.WithSourceCode(randomSourceCode()),
	)
	id, err := s.Create()
	if err != nil {
		t.Fatalf("创建提交失败: %v", err)
	}
	s.Id = id
	s.Score = 80
	err = s.Update()
	if err != nil {
		t.Fatalf("更新提交失败: %v", err)
	}
}

// 测试更新不存在的提交
func TestSubmissionUpdate_NotFound(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithId(99999999),
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(20),
		submission.WithLanguageId(1),
		submission.WithSourceCode(randomSourceCode()),
	)
	err := s.Update()
	if err == nil {
		t.Fatalf("更新不存在的提交应失败")
	}
}

// 测试提交删除成功
func TestSubmissionDelete_Success(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(20),
		submission.WithLanguageId(1),
		submission.WithSourceCode(randomSourceCode()),
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
func TestSubmissionDelete_NotFound(t *testing.T) {
	s := submission.NewSubmission(
		submission.WithId(99999999),
		submission.WithUserId(1),
		submission.WithProblemId(1),
		submission.WithStatus(entity.JudgeAC),
		submission.WithScore(100),
		submission.WithMemory(128),
		submission.WithTime(0.123),
		submission.WithLength(20),
		submission.WithLanguageId(1),
		submission.WithSourceCode(randomSourceCode()),
	)
	err := s.Delete()
	if err == nil {
		t.Fatalf("删除不存在的提交应失败")
	}
}
