package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"

	"gorm.io/gorm"
)

func InsertCollectionUser(cu entity.CollectionUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionUser{}).Create(&cu).Error
	})
}

func DeleteCollectionUser(cu entity.CollectionUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionUser{}).Delete(&cu).Error
	})
}
