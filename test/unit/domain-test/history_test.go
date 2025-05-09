package domain_test

import (
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/history/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"testing"
	"time"
)

// 生成随机标题
func randomHistoryTitle() string {
	return "历史标题" + time.Now().Format("150405.000")
}

// 生成随机描述
func randomHistoryDescription() string {
	return "历史描述" + time.Now().Format("150405.000")
}

// 生成随机输入输出
func randomInput() string {
	return "输入" + time.Now().Format("150405.000")
}
func randomOutput() string {
	return "输出" + time.Now().Format("150405.000")
}
func randomSource() string {
	return "来源" + time.Now().Format("150405.000")
}

// 测试历史记录创建成功
func TestHistoryCreate_Success(t *testing.T) {
	h := history.NewHistory(
		history.WithUserId(1),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	id, err := h.Create()
	if err != nil {
		t.Fatalf("创建历史记录失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建历史记录返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestHistoryCreate_EmptyUserId(t *testing.T) {
	h := history.NewHistory(
		history.WithUserId(0),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	_, err := h.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试历史记录更新成功
func TestHistoryUpdate_Success(t *testing.T) {
	h := history.NewHistory(
		history.WithUserId(1),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	id, err := h.Create()
	if err != nil {
		t.Fatalf("创建历史记录失败: %v", err)
	}
	h.Id = id
	h.Description = valueobject.NewDescription("更新后的描述")
	err = h.Update()
	if err != nil {
		t.Fatalf("更新历史记录失败: %v", err)
	}
}

// 测试更新不存在的历史记录
func TestHistoryUpdate_NotFound(t *testing.T) {
	h := history.NewHistory(
		history.WithId(99999999),
		history.WithUserId(1),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	err := h.Update()
	if err == nil {
		t.Fatalf("更新不存在的历史记录应失败")
	}
}

// 测试历史记录删除成功
func TestHistoryDelete_Success(t *testing.T) {
	h := history.NewHistory(
		history.WithUserId(1),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	id, err := h.Create()
	if err != nil {
		t.Fatalf("创建历史记录失败: %v", err)
	}
	h.Id = id
	err = h.Delete()
	if err != nil {
		t.Fatalf("删除历史记录失败: %v", err)
	}
}

// 测试删除不存在的历史记录
func TestHistoryDelete_NotFound(t *testing.T) {
	h := history.NewHistory(
		history.WithId(99999999),
		history.WithUserId(1),
		history.WithProblemId(1),
		history.WithTitle(randomHistoryTitle()),
		history.WithSource(randomSource()),
		history.WithDifficulty(entity.DifficultyE),
		history.WithTimeLimit(1.0),
		history.WithMemoryLimit(128),
		history.WithDescription(randomHistoryDescription()),
		history.WithInput(randomInput()),
		history.WithOutput(randomOutput()),
		history.WithSampleInput(randomInput()),
		history.WithSampleOutput(randomOutput()),
		history.WithHint(randomHistoryDescription()),
		history.WithOperation(entity.OperationInsert),
	)
	err := h.Delete()
	if err == nil {
		t.Fatalf("删除不存在的历史记录应失败")
	}
}
