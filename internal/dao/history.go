package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
)

type auxiliaryHistory struct {
	entity.History
	model.BriefUser
}

// 插入题目历史记录
func InsertHistory(ph entity.History) (uint64, error) {
	tx := db.Db.Create(&ph)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return ph.Id, nil
}

// 根据题目ID查询题目历史记录
func SelectHistoriesByProblemId(pid uint64) ([]entity.History, error) {
	var auxiliaryHistorys []auxiliaryHistory
	var phs []entity.History

	condition := model.HistoryWhere{}

	tx := db.Db.Model(&entity.History{}).Where("tbl_history.problem_id = ?", pid)
	tx = condition.GenerateWhere()(tx)
	tx = tx.Find(&auxiliaryHistorys)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, auxiliaryHistory := range auxiliaryHistorys {
		ph := auxiliaryHistory.History
		ph.User = entity.User{
			Id:       auxiliaryHistory.UserId,
			Username: auxiliaryHistory.Username,
			Role:     auxiliaryHistory.Role,
			Avatar:   auxiliaryHistory.Avatar,
		}
		phs = append(phs, ph)
	}

	return phs, nil
}

// 根据创建时间统计用户数量
func CountHistoriesBetweenCreateTimeByOperation(startTime time.Time, endTime time.Time, operation entity.Operation) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.History{}).Where("create_time between ? and ? AND operation = ?", startTime, endTime, operation).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
