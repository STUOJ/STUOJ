package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
)

type historyStore struct{}

var HistoryStore = new(historyStore)

// 插入题目历史记录
func (store *historyStore) Insert(ph entity.History) (uint64, error) {
	tx := db.Db.Create(&ph)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return ph.Id, nil
}

// 根据题目ID查询题目历史记录
func (store *historyStore) SelectByProblemId(pid uint64) ([]entity.History, error) {
	var phs []entity.History

	condition := model.HistoryWhere{}

	tx := db.Db.Model(&entity.History{}).Where("tbl_history.problem_id = ?", pid)
	tx = condition.GenerateWhere()(tx)
	tx = tx.Find(&phs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return phs, nil
}
