package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"

	"gorm.io/gorm"
)

type teamStore struct{}

var TeamStore = new(teamStore)

func (store *teamStore) Insert(c entity.Team) (uint64, error) {
	err := db.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Team{}).Create(&c).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}

func (store *teamStore) SelectById(id uint64) (entity.Team, error) {
	var c entity.Team

	condition := model.TeamWhere{}
	condition.Id.Set(id)

	tx := db.Db.Model(&entity.Team{})
	where := condition.GenerateWhere()
	tx = tx.Where(&entity.Team{Id: id})
	tx = where(tx)
	tx = tx.Scan(&c)

	if tx.Error != nil {
		return entity.Team{}, tx.Error
	}

	return c, nil
}

func (store *teamStore) Select(condition model.TeamWhere) ([]entity.Team, error) {
	var teams []entity.Team
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Team{})
	tx = where(tx)
	tx = tx.Scan(&teams)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return teams, nil
}

func (store *teamStore) UpdateById(c entity.Team) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func (store *teamStore) DeleteById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Team{}).Error
	})
}

func (store *teamStore) Count(condition model.TeamWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Team{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

func (store *teamStore) InsertTeamUser(tu entity.TeamUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.TeamUser{}).Create(&tu).Error
	})
}

func (store *teamStore) DeleteTeamUser(tu entity.TeamUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.TeamUser{}).Delete(&tu).Error
	})
}
