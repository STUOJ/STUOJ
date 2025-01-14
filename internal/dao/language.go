package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
)

// 插入语言
func InsertLanguage(l entity.Language) (uint64, error) {
	tx := db.Db.Create(&l)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return l.Id, nil
}

// 查询语言
func SelectLanguage(con model.LanguageWhere) ([]entity.Language, error) {
	var languages []entity.Language
	where := con.GenerateWhere()
	tx := db.Db.Model(&entity.Language{})
	tx = where(tx)
	tx = tx.Find(&languages)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return languages, nil
}

// 根据ID查询标签
func SelectLanguageById(id uint64) (entity.Language, error) {
	var l entity.Language
	tx := db.Db.Where("id = ?", id).First(&l)
	if tx.Error != nil {
		return entity.Language{}, tx.Error
	}

	return l, nil
}

// 根据名字模糊查询语言
func SelectLanguageLikeName(name string) (entity.Language, error) {
	var l entity.Language

	tx := db.Db.Where("name like ?", "%"+name+"%").First(&l)
	if tx.Error != nil {
		return entity.Language{}, tx.Error
	}

	return l, nil
}

func UpdateLanguage(l entity.Language) error {
	tx := db.Db.Model(&l).Updates(l)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 删除所有语言
func DeleteAllLanguages() error {
	tx := db.Db.Where("1 = 1").Delete(&entity.Language{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计语言数量
func CountLanguages() (uint64, error) {
	var count int64

	tx := db.Db.Model(&entity.Language{}).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
