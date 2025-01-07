package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type CommentWhere struct {
	UserId    model.Field[uint64]
	BlogId    model.Field[uint64]
	Status    model.Field[entity.CommentStatus]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.Field[uint64]
	Size      model.Field[uint64]
}

// 根据ID查询评论
func SelectCommentById(id uint64) (entity.Comment, error) {
	var c entity.Comment

	tx := db.Db.Where("id = ?", id).First(&c)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}

	return c, nil
}

// 查询评论
func SelectComments(condition CommentWhere) ([]entity.Comment, error) {
	var comments []entity.Comment
	where := generateCommentWhereCondition(condition)
	tx := db.Db.Model(&entity.Comment{})
	tx = where(tx)
	tx = tx.Find(&comments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return comments, nil
}

// 插入评论
func InsertComment(c entity.Comment) (uint64, error) {
	tx := db.Db.Create(&c)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return c.Id, nil
}

// 根据ID更新评论
func UpdateCommentById(b entity.Comment) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除评论
func DeleteCommentById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Comment{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计评论数量
func CountComments(condition CommentWhere) (uint64, error) {
	var count int64
	where := generateCommentWhereConditionWithNoPage(condition)

	tx := db.Db.Model(&entity.Comment{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 根据创建时间统计博客数量
func CountCommentsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var counts []model.CountByDate

	tx := db.Db.Model(&entity.Comment{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

func generateCommentWhereConditionWithNoPage(con CommentWhere) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.UserId.Exist() {
			whereClause["user_id"] = con.UserId.Value()
		}
		if con.BlogId.Exist() {
			whereClause["blog_id"] = con.BlogId.Value()
		}
		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		where := db.Where(whereClause)
		if con.StartTime.Exist() {
			where.Where("create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where.Where("create_time <= ?", con.EndTime.Value())
		}
		return where
	}
}

func generateCommentWhereCondition(con CommentWhere) func(*gorm.DB) *gorm.DB {
	where := generateCommentWhereConditionWithNoPage(con)
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
