package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入题目
func Insert(p entity.Problem, uid uint64) (uint64, error) {
	var err error

	updateTime := time.Now()
	p.UpdateTime = updateTime
	p.CreateTime = updateTime

	// 插入题目
	p.Id, err = dao.InsertProblem(p)
	if err != nil {
		return 0, errors.New("插入题目失败")
	}

	ph := entity.History{
		UserId:       uid,
		ProblemId:    p.Id,
		Title:        p.Title,
		Source:       p.Source,
		Difficulty:   p.Difficulty,
		TimeLimit:    p.TimeLimit,
		MemoryLimit:  p.MemoryLimit,
		Description:  p.Description,
		Input:        p.Input,
		Output:       p.Output,
		SampleInput:  p.SampleInput,
		SampleOutput: p.SampleOutput,
		Hint:         p.Hint,
		Operation:    entity.OperationInsert,
		CreateTime:   updateTime,
	}

	// 添加题目历史记录
	ph.Id, err = dao.InsertHistory(ph)
	if err != nil {
		log.Println(err)
		return p.Id, errors.New("插入题目成功，但插入题目历史记录失败")
	}

	return p.Id, nil
}

// 给题目添加标签
func InsertTag(pid uint64, tid uint64, uid uint64, role entity.Role) error {
	// 初始化题目标签
	pt := entity.ProblemTag{
		ProblemId: pid,
		TagId:     tid,
	}

	// 读取题目
	_, uids, err := dao.SelectProblemByIdWithUser(pid)
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	if role < entity.RoleAdmin {
		userIdsMap := make(map[uint64]struct{})
		for _, uid := range uids {
			userIdsMap[uid] = struct{}{}
		}
		if _, exists := userIdsMap[uid]; !exists {
			return errors.New("没有该题权限")
		}
	}

	// 读取标签
	_, err = dao.SelectTagById(tid)
	if err != nil {
		log.Println(err)
		return errors.New("标签不存在")
	}

	// 检查题目标签关系是否存在
	count, err := dao.CountProblemTag(pt)
	if err != nil || count > 0 {
		if err != nil {
			log.Println(err)
		}
		return errors.New("该题目已存在该标签")
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 插入题目标签
	err = dao.InsertProblemTag(pt)
	if err != nil {
		log.Println(err)
		return errors.New("添加失败，该题目已存在该标签")
	}

	return nil
}

// 插入题目历史记录
func InsertHistory(p entity.Problem, uid uint64, op entity.Operation) (uint64, error) {
	var err error

	updateTime := time.Now()
	ph := entity.History{
		UserId:       uid,
		ProblemId:    p.Id,
		Title:        p.Title,
		Source:       p.Source,
		Difficulty:   p.Difficulty,
		TimeLimit:    p.TimeLimit,
		MemoryLimit:  p.MemoryLimit,
		Description:  p.Description,
		Input:        p.Input,
		Output:       p.Output,
		SampleInput:  p.SampleInput,
		SampleOutput: p.SampleOutput,
		Hint:         p.Hint,
		Operation:    op,
		CreateTime:   updateTime,
	}

	ph.Id, err = dao.InsertHistory(ph)
	if err != nil {
		return 0, err
	}

	return ph.Id, nil
}
