package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type auxiliaryCollection struct {
	entity.Collection
	model.BriefUser
	CollectionUserIds    string `gorm:"column:collection_user_id"`
	CollectionProblemIds string `gorm:"column:collection_problem_id"`
}

func InsertCollection(c entity.Collection) (uint64, error) {
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

func SelectCollectionById(id uint64) (entity.Collection, error) {
	var c auxiliaryCollection

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
	userIds := make([]uint64, 0)
	if c.CollectionUserIds != "" {
		for _, idStr := range strings.Split(c.CollectionUserIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				userIds = append(userIds, id)
			}
		}
	}
	c.Collection.UserIds = userIds

	problemIds := make([]uint64, 0)
	if c.CollectionProblemIds != "" {
		for _, idStr := range strings.Split(c.CollectionProblemIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				problemIds = append(problemIds, id)
			}
		}
	}
	c.Collection.ProblemIds = problemIds
	c.Collection.User = entity.User{
		Id:       c.UserId,
		Username: c.Username,
		Role:     c.Role,
		Avatar:   c.Avatar,
	}

	return c.Collection, nil
}

func SelectCollections(condition model.CollectionWhere) ([]entity.Collection, error) {
	var auxiliaryCollections []auxiliaryCollection
	var collections []entity.Collection
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Collection{})
	tx = where(tx)
	tx = tx.Scan(&auxiliaryCollections)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for i := range auxiliaryCollections {
		userIds := make([]uint64, 0)
		problemIds := make([]uint64, 0)
		if auxiliaryCollections[i].CollectionUserIds != "" {
			for _, idStr := range strings.Split(auxiliaryCollections[i].CollectionUserIds, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					userIds = append(userIds, id)
				}
			}
		}
		auxiliaryCollections[i].Collection.UserIds = userIds
		if auxiliaryCollections[i].CollectionProblemIds != "" {
			for _, idStr := range strings.Split(auxiliaryCollections[i].CollectionProblemIds, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					problemIds = append(problemIds, id)
				}
			}
			auxiliaryCollections[i].Collection.ProblemIds = problemIds
		}
		auxiliaryCollections[i].ProblemIds = problemIds
		auxiliaryCollections[i].Collection.User = entity.User{
			Id:       auxiliaryCollections[i].UserId,
			Username: auxiliaryCollections[i].Username,
			Role:     auxiliaryCollections[i].Role,
			Avatar:   auxiliaryCollections[i].Avatar,
		}
		collections = append(collections, auxiliaryCollections[i].Collection)
	}

	return collections, nil
}

func UpdateCollectionById(c entity.Collection) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&c).Updates(c).Error
	})
}

func DeleteCollectionById(id uint64) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).Delete(&entity.Collection{}).Error
	})
}

func CountCollections(condition model.CollectionWhere) (uint64, error) {
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

func InsertCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Create(&cp).Error
	})
}

func UpdateCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Where(&entity.CollectionProblem{CollectionId: cp.CollectionId, ProblemId: cp.ProblemId}).Updates(&cp).Error
	})
}

func DeleteCollectionProblem(cp entity.CollectionProblem) error {
	return db.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&entity.CollectionProblem{}).Delete(&cp).Error
	})
}

// 根据ID更新题单更新时间
func UpdateCollectionUpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Collection{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
