package domain_test

import (
	"STUOJ/internal/domain/contest"
	"STUOJ/internal/domain/contest/valueobject"
	"STUOJ/internal/infrastructure/persistence/entity"
	"testing"
	"time"
)

// 生成随机标题
func randomContestTitle() string {
	return "比赛标题" + time.Now().Format("150405.000")
}

// 生成随机描述
func randomContestDescription() string {
	return "比赛描述" + time.Now().Format("150405.000")
}

// 测试比赛创建成功
func TestContestCreate_Success(t *testing.T) {
	c := contest.NewContest(
		contest.WithUserId(1),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建比赛失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建比赛返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestContestCreate_EmptyUserId(t *testing.T) {
	c := contest.NewContest(
		contest.WithUserId(0),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	_, err := c.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试比赛更新成功
func TestContestUpdate_Success(t *testing.T) {
	c := contest.NewContest(
		contest.WithUserId(1),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建比赛失败: %v", err)
	}
	c.Id.Set(id)
	c.Description = valueobject.NewDescription("更新后的描述")
	err = c.Update()
	if err != nil {
		t.Fatalf("更新比赛失败: %v", err)
	}
}

// 测试更新不存在的比赛
func TestContestUpdate_NotFound(t *testing.T) {
	c := contest.NewContest(
		contest.WithId(99999999),
		contest.WithUserId(1),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	err := c.Update()
	if err == nil {
		t.Fatalf("更新不存在的比赛应失败")
	}
}

// 测试比赛删除成功
func TestContestDelete_Success(t *testing.T) {
	c := contest.NewContest(
		contest.WithUserId(1),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	id, err := c.Create()
	if err != nil {
		t.Fatalf("创建比赛失败: %v", err)
	}
	c.Id.Set(id)
	err = c.Delete()
	if err != nil {
		t.Fatalf("删除比赛失败: %v", err)
	}
}

// 测试删除不存在的比赛
func TestContestDelete_NotFound(t *testing.T) {
	c := contest.NewContest(
		contest.WithId(99999999),
		contest.WithUserId(1),
		contest.WithTitle(randomContestTitle()),
		contest.WithDescription(randomContestDescription()),
		contest.WithStatus(entity.ContestPublic),
		contest.WithFormat(entity.ContestACM),
		contest.WithTeamSize(1),
		contest.WithStartTime(time.Now().Add(time.Hour)),
		contest.WithEndTime(time.Now().Add(2*time.Hour)),
	)
	err := c.Delete()
	if err == nil {
		t.Fatalf("删除不存在的比赛应失败")
	}
}
