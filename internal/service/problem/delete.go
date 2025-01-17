package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 根据ID删除题目
func DeleteByProblemId(pid uint64, uid uint64) error {
	// 查询题目是否存在
	_, err := dao.SelectProblemById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("题目不存在")
	}

	// 删除题目的所有标签
	err = dao.DeleteProblemTagsByProblemId(pid)
	if err != nil {
		log.Println(err)
		return errors.New("删除题目标签失败")
	}

	// 删除题目的所有评测点
	err = dao.DeleteTestcasesByProblemId(pid)
	if err != nil {
		log.Println(err)
		return errors.New("删除评测点失败")
	}

	// 添加题目历史记录
	updateTime := time.Now()
	ph := entity.History{
		UserId:     uid,
		ProblemId:  pid,
		Operation:  entity.OperationDelete,
		CreateTime: updateTime,
	}
	_, err = dao.InsertHistory(ph)
	if err != nil {
		log.Println(err)
		return errors.New("插入题目历史记录失败")
	}

	// 删除题目
	err = dao.DeleteProblemById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("删除题目失败")
	}

	return nil
}

// 删除题目的某个标签
func DeleteTag(pid uint64, tid uint64, uid uint64, role entity.Role) error {
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
	if err != nil || count < 1 {
		if err != nil {
			log.Println(err)
		}
		return errors.New("该题目不存在该标签")
	}

	// 更新题目更新时间
	err = dao.UpdateProblemUpdateTimeById(pid)
	if err != nil {
		log.Println(err)
		return errors.New("更新题目更新时间失败")
	}

	// 删除题目标签
	err = dao.DeleteProblemTag(pt)
	if err != nil {
		log.Println(err)
		return errors.New("删除失败")
	}

	return nil
}
