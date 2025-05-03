package domain_test

import (
	"STUOJ/internal/domain/team"
	"STUOJ/internal/domain/team/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
	"math/rand"
	"strconv"
	"testing"
)

// 生成随机队伍名和描述
func randomTeamName() valueobject.Name {
	return valueobject.NewName("队伍" + strconv.Itoa(rand.Intn(100000)))
}
func randomTeamDesc() valueobject.Description {
	return valueobject.NewDescription("描述" + strconv.Itoa(rand.Intn(100000)))
}

// 测试队伍创建成功
func TestTeamCreate_Success(t *testing.T) {
	teamObj := &team.Team{
		UserId:      1,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	id, err := teamObj.Create()
	if err != nil {
		t.Fatalf("创建队伍失败: %v", err)
	}
	if id == 0 {
		t.Fatalf("创建队伍返回的ID无效")
	}
}

// 测试用户ID为空时创建失败
func TestTeamCreate_EmptyUserId(t *testing.T) {
	teamObj := &team.Team{
		UserId:      0,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	_, err := teamObj.Create()
	if err == nil {
		t.Fatalf("用户ID为空时应创建失败")
	}
}

// 测试队伍更新成功
func TestTeamUpdate_Success(t *testing.T) {
	teamObj := &team.Team{
		UserId:      1,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	id, err := teamObj.Create()
	if err != nil {
		t.Fatalf("创建队伍失败: %v", err)
	}
	teamObj.Id = id
	teamObj.Description = valueobject.NewDescription("更新后的描述")
	err = teamObj.Update()
	if err != nil {
		t.Fatalf("更新队伍失败: %v", err)
	}
}

// 测试更新不存在的队伍
func TestTeamUpdate_NotFound(t *testing.T) {
	teamObj := &team.Team{
		Id:          99999999,
		UserId:      1,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	err := teamObj.Update()
	if err == nil {
		t.Fatalf("更新不存在的队伍应失败")
	}
}

// 测试队伍删除成功
func TestTeamDelete_Success(t *testing.T) {
	teamObj := &team.Team{
		UserId:      1,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	id, err := teamObj.Create()
	if err != nil {
		t.Fatalf("创建队伍失败: %v", err)
	}
	teamObj.Id = id
	err = teamObj.Delete()
	if err != nil {
		t.Fatalf("删除队伍失败: %v", err)
	}
}

// 测试删除不存在的队伍
func TestTeamDelete_NotFound(t *testing.T) {
	teamObj := &team.Team{
		Id:          99999999,
		UserId:      1,
		ContestId:   1,
		Name:        randomTeamName(),
		Description: randomTeamDesc(),
		Status:      entity.TeamEnabled,
	}
	err := teamObj.Delete()
	if err == nil {
		t.Fatalf("删除不存在的队伍应失败")
	}
}
