package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
)

// 插入标签
func InsertTag(t entity.Tag) (uint64, error) {
	tx := db.Db.Create(&t)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return t.Id, nil
}

// 根据ID查询标签
func SelectTagById(id uint64) (entity.Tag, error) {
	var t entity.Tag
	tx := db.Db.Where("id = ?", id).First(&t)
	if tx.Error != nil {
		return entity.Tag{}, tx.Error
	}

	return t, nil
}

// 查询所有标签
func SelectAllTags() ([]entity.Tag, error) {
	var tags []entity.Tag
	tx := db.Db.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 查询标签
func SelectTags(condition model.TagWhere) ([]entity.Tag, error) {
	var tags []entity.Tag
	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.Tag{})
	tx = where(tx)
	tx = tx.Find(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tags, nil
}

// 根据ID更新标签
func UpdateTagById(t entity.Tag) error {
	tx := db.Db.Model(&t).Where("id = ?", t.Id).Updates(t)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除标签
func DeleteTagById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Tag{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计标签数量
func CountTags(condition model.TagWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.Tag{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
