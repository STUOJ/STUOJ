package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
)

type commentStore struct{}

var CommentStore = new(commentStore)

// 根据ID查询评论
func (store *commentStore) SelectById(id uint64) (entity.Comment, error) {
	var c entity.Comment

	condition := model.CommentWhere{}

	tx := db.Db.Where(&entity.Comment{Id: id})
	tx = condition.GenerateWhere()(tx)
	tx = tx.First(&c)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	return c, nil
}

// 查询评论
func (store *commentStore) Select(condition model.CommentWhere) ([]entity.Comment, error) {
	var comments []entity.Comment
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Comment{})
	tx = where(tx)
	tx = tx.Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 插入评论
func (store *commentStore) Insert(c entity.Comment) (uint64, error) {
	tx := db.Db.Create(&c)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return c.Id, nil
}

// 根据ID更新评论
func (store *commentStore) UpdateById(b entity.Comment) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除评论
func (store *commentStore) DeleteById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Comment{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计评论数量
func (store *commentStore) Count(condition model.CommentWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()

	tx := db.Db.Model(&entity.Comment{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
