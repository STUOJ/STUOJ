package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
	"time"
)

type problemStore struct{}

var ProblemStore = new(problemStore)

// 插入题目
func (store *problemStore) Insert(p entity.Problem) (uint64, error) {
	tx := db.Db.Create(&p)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return p.Id, nil
}

func (store *problemStore) SelectById(id uint64) (entity.Problem, error) {
	var p entity.Problem

	tx := db.Db.Model(&entity.Problem{})
	tx = tx.Where(&entity.Problem{Id: id}).
		Scan(&p)

	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	return p, nil
}

func (store *problemStore) Select(condition model.ProblemWhere) ([]entity.Problem, error) {
	var problems []entity.Problem

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.Scan(&problems)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return problems, nil
}

// 根据ID更新题目
func (store *problemStore) UpdateById(p entity.Problem) error {
	// 明确指定要更新的字段，包含需要处理空值的字段
	tx := db.Db.Model(&p).
		Select("title", "source", "difficulty", "time_limit", "memory_limit", "description", "input", "output", "sample_input", "sample_output", "hint", "status", "update_time").
		Updates(p)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID更新题目更新时间
func (store *problemStore) UpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Problem{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除题目
func (store *problemStore) DeleteById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Problem{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计题目数量
func (store *problemStore) Count(condition model.ProblemWhere) (uint64, error) {
	var count int64

	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
