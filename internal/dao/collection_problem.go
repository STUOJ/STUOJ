package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"

	"gorm.io/gorm"
)

func InsertCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Create(&cp).Error
	})
}

func DeleteCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Delete(&cp).Error
	})
}
