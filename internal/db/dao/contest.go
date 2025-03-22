package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"

	"gorm.io/gorm"
)

type contestStore struct{}

var ContestStore = new(contestStore)

func (store *contestStore) Insert(c entity.Contest) (uint64, error) {
	err := db.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Contest{}).Create(&c).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}

// 根据ID查询比赛
func (store *contestStore) SelectById(id uint64) (entity.Contest, error) {
	var ct entity.Contest
	condition := model.ContestWhere{}
	tx := db.Db.Model(&entity.Contest{})
	tx = tx.Where(&entity.Contest{Id: id})
	tx = condition.GenerateWhere()(tx)
	tx = tx.Find(&ct)
	if tx.Error != nil {
		return entity.Contest{}, tx.Error
	}

	return ct, nil
}

func (store *contestStore) Select(condition model.ContestWhere) ([]entity.Contest, error) {
	var contests []entity.Contest

	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.Contest{})
	tx = where(tx)
	tx = tx.Find(&contests)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return contests, nil
}

func (store *contestStore) UpdateById(c entity.Contest) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func (store *contestStore) DeleteById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Contest{}).Error
	})
}

func (store *contestStore) Count(condition model.ContestWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Contest{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
