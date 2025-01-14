package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type auxiliaryComment struct {
	entity.Comment
	BriefUser
	BriefBlog
}

// 根据ID查询评论
func SelectCommentById(id uint64) (entity.Comment, error) {
	var auxiliaryComment auxiliaryComment
	var c entity.Comment

	tx := db.Db.Where(&entity.Comment{Id: id})
	tx = commentUnionJoins(tx)
	tx = tx.First(&auxiliaryComment)
	if tx.Error != nil {
		return entity.Comment{}, tx.Error
	}
	c = auxiliaryComment.Comment
	c.User = entity.User{
		Id:       auxiliaryComment.UserId,
		Username: auxiliaryComment.Username,
		Role:     auxiliaryComment.Role,
		Avatar:   auxiliaryComment.Avatar,
	}
	c.Blog = entity.Blog{
		Id:     auxiliaryComment.BlogId,
		Title:  auxiliaryComment.BlogTitle,
		Status: auxiliaryComment.BlogStatus,
	}

	return c, nil
}

// 查询评论
func SelectComments(condition model.CommentWhere) ([]entity.Comment, error) {
	var auxiliaryComments []auxiliaryComment
	var comments []entity.Comment
	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Comment{})
	tx = where(tx)
	tx = commentUnionJoins(tx)
	tx = tx.Find(&auxiliaryComments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, auxiliaryComment := range auxiliaryComments {
		comment := auxiliaryComment.Comment
		comment.User = entity.User{
			Id:       auxiliaryComment.UserId,
			Username: auxiliaryComment.Username,
			Role:     auxiliaryComment.Role,
			Avatar:   auxiliaryComment.Avatar,
		}
		comment.Blog = entity.Blog{
			Id:     auxiliaryComment.BlogId,
			Title:  auxiliaryComment.BlogTitle,
			Status: auxiliaryComment.BlogStatus,
		}
		comments = append(comments, comment)
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
func CountComments(condition model.CommentWhere) (uint64, error) {
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

// 根据创建时间统计博客数量
func CountCommentsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var counts []model.CountByDate

	tx := db.Db.Model(&entity.Comment{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

func commentUnionJoins(tx *gorm.DB) *gorm.DB {
	query := []string{"tbl_comment.*"}
	query = append(query, briefUserSelect()...)
	query = append(query, briefBlogSelect()...)
	tx = tx.Select(query)
	tx = briefUserJoins(tx, "tbl_comment")
	tx = briefBlogJoins(tx, "tbl_comment")
	return tx
}
