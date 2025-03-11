package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"gorm.io/gorm"
)

func InsertContest(c entity.Contest) (uint64, error) {
	err := db.Db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&entity.Contest{}).Create(&c).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}
