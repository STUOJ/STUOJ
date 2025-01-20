package history

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
)

// 根据题目ID查询历史记录
func SelectByProblemId(pid uint64) ([]entity.History, error) {
	histories, err := dao.SelectHistoriesByProblemId(pid)
	if err != nil {
		log.Println(err)
		return nil, errors.New("获取题目历史记录失败")
	}

	return histories, nil
}
