package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID更新题目
func UpdateById(p entity.Problem, uid uint64, role entity.Role) error {
	// 读取题目
	p0, userIds, err := dao.SelectProblemByIdWithUser(p.Id)
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	if role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range userIds {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[uid]; !exists {
			return errors.New("没有该题权限")
		}
	}

	updateTime := time.Now()

	// 添加题目历史记录
	ph := entity.History{
		UserId:       uid,
		ProblemId:    p0.Id,
		Title:        p0.Title,
		Source:       p0.Source,
		Difficulty:   p0.Difficulty,
		TimeLimit:    p0.TimeLimit,
		MemoryLimit:  p0.MemoryLimit,
		Description:  p0.Description,
		Input:        p0.Input,
		Output:       p0.Output,
		SampleInput:  p0.SampleInput,
		SampleOutput: p0.SampleOutput,
		Hint:         p0.Hint,
		Operation:    entity.OperationUpdate,
		CreateTime:   updateTime,
	}
	ph.Id, err = dao.InsertHistory(ph)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目成功，但插入题目历史记录失败")
	}

	// 更新题目
	p.UpdateTime = updateTime
	err = dao.UpdateProblemById(p)
	if err != nil {
		return errors.New("更新题目失败")
	}

	return nil
}
