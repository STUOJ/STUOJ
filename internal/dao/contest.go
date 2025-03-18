package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"gorm.io/gorm"
)

type auxiliaryContest struct {
	entity.Contest
	model.BriefUser
}

func InsertContest(c entity.Contest) (uint64, error) {
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

func SelectContestById(id uint64) (entity.Contest, error) {
	var c auxiliaryContest

	tx := db.Db.Model(&entity.Contest{})
	tx = tx.Where(&entity.Contest{Id: id}).
		Scan(&c)

	if tx.Error != nil {
		return entity.Contest{}, tx.Error
	}

	c.Contest.User = entity.User{
		Id:       c.UserId,
		Username: c.Username,
		Role:     c.Role,
		Avatar:   c.Avatar,
	}

	return c.Contest, nil
}

func SelectContests(condition model.ContestWhere) ([]entity.Contest, error) {
	var auxiliaryContests []auxiliaryContest
	var contests []entity.Contest
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Contest{})
	tx = where(tx)
	tx = tx.Scan(&auxiliaryContests)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return contests, nil
}

func UpdateContestById(c entity.Contest) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func DeleteContestById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Contest{}).Error
	})
}

func CountContests(condition model.ContestWhere) (uint64, error) {
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
