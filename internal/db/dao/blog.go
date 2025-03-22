package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
)

type blogStore struct{}

var BlogStore = new(blogStore)

// 插入博客
func (store *blogStore) Insert(b entity.Blog) (uint64, error) {
	tx := db.Db.Create(&b)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return b.Id, nil
}

// 根据ID查询博客
func (store *blogStore) SelectById(id uint64) (entity.Blog, error) {
	var b entity.Blog
	condition := model.BlogWhere{}
	tx := db.Db.Model(&entity.Blog{})
	tx = tx.Where(&entity.Blog{Id: id})
	tx = condition.GenerateWhere()(tx)
	tx = tx.Find(&b)
	if tx.Error != nil {
		return entity.Blog{}, tx.Error
	}

	return b, nil
}

func (store *blogStore) Select(condition model.BlogWhere) ([]entity.Blog, error) {
	var blogs []entity.Blog

	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.Blog{})
	tx = where(tx)
	tx = tx.Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return blogs, nil
}

// 根据ID更新博客
func (store *blogStore) UpdateById(b entity.Blog) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除博客
func (store *blogStore) DeleteById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Blog{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计博客数量
func (store *blogStore) Count(condition model.BlogWhere) (uint64, error) {
	var count int64

	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.Blog{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
