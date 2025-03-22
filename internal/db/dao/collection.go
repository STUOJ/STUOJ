package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type collectionStore struct{}

var CollectionStore = new(collectionStore)

func (store *collectionStore) Insert(c entity.Collection) (uint64, error) {
	err := db.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Collection{}).Create(&c).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}

func (store *collectionStore) SelectById(id uint64) (entity.Collection, error) {
	var c entity.Collection

	condition := model.CollectionWhere{}
	condition.Id.Set(id)

	tx := db.Db.Model(&entity.Collection{})
	where := condition.GenerateWhere()
	tx = tx.Where(&entity.Collection{Id: id})
	tx = where(tx)
	tx = tx.Scan(&c)

	if tx.Error != nil {
		return entity.Collection{}, tx.Error
	}

	return c, nil
}

func (store *collectionStore) Select(condition model.CollectionWhere) ([]entity.Collection, error) {
	var collections []entity.Collection
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Collection{})
	tx = where(tx)
	tx = tx.Scan(&collections)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return collections, nil
}

func (store *collectionStore) UpdateById(c entity.Collection) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func (store *collectionStore) DeleteById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Collection{}).Error
	})
}

func (store *collectionStore) Count(condition model.CollectionWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Collection{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

func (store *collectionStore) InsertCollectionUser(cu entity.CollectionUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionUser{}).Create(&cu).Error
	})
}

func (store *collectionStore) DeleteCollectionUser(cu entity.CollectionUser) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionUser{}).Delete(&cu).Error
	})
}

func (store *collectionStore) InsertCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Create(&cp).Error
	})
}

func (store *collectionStore) UpdateCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Where(&entity.CollectionProblem{CollectionId: cp.CollectionId, ProblemId: cp.ProblemId}).Updates(&cp).Error
	})
}

func (store *collectionStore) DeleteCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Delete(&cp).Error
	})
}

// 根据ID更新题单更新时间
func (store *collectionStore) UpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Collection{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
