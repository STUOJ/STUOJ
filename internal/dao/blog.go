package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BriefBlog struct {
	BlogTitle  string            `gorm:"column:blog_title"`
	BlogStatus entity.BlogStatus `gorm:"column:blog_status"`
}

type auxiliaryBlog struct {
	entity.Blog
	BriefUser
	BriefProblem
}

// 插入博客
func InsertBlog(b entity.Blog) (uint64, error) {
	tx := db.Db.Create(&b)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return b.Id, nil
}

// 根据ID查询博客
func SelectBlogById(id uint64) (entity.Blog, error) {
	var auxiliaryBlog auxiliaryBlog
	var b entity.Blog
	tx := db.Db.Model(&entity.Blog{})
	tx = tx.Where(&entity.Blog{Id: id})
	tx = blogUnionJoins(tx)
	tx = tx.Find(&auxiliaryBlog)
	if tx.Error != nil {
		return entity.Blog{}, tx.Error
	}
	b = auxiliaryBlog.Blog
	b.User = entity.User{
		Id:       auxiliaryBlog.UserId,
		Username: auxiliaryBlog.Username,
		Role:     auxiliaryBlog.Role,
		Avatar:   auxiliaryBlog.Avatar,
	}
	if auxiliaryBlog.ProblemId != 0 {
		b.Problem = entity.Problem{
			Id:         auxiliaryBlog.ProblemId,
			Title:      auxiliaryBlog.ProblemTitle,
			Status:     auxiliaryBlog.ProblemStatus,
			Difficulty: auxiliaryBlog.ProblemDifficulty,
		}
	}

	return b, nil
}

func SelectBlogs(condition model.BlogWhere) ([]entity.Blog, error) {
	var auxiliaryBlogs []auxiliaryBlog
	var blogs []entity.Blog

	where := condition.GenerateWhere()

	tx := db.Db.Model(&entity.Blog{})
	tx = where(tx)
	tx = blogUnionJoins(tx)
	tx = tx.Find(&auxiliaryBlogs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	for _, auxiliaryBlog := range auxiliaryBlogs {
		blog := auxiliaryBlog.Blog
		blog.User = entity.User{
			Id:       auxiliaryBlog.UserId,
			Username: auxiliaryBlog.Username,
			Role:     auxiliaryBlog.Role,
			Avatar:   auxiliaryBlog.Avatar,
		}
		if auxiliaryBlog.ProblemId != 0 {
			blog.Problem = entity.Problem{
				Id:         auxiliaryBlog.ProblemId,
				Title:      auxiliaryBlog.ProblemTitle,
				Status:     auxiliaryBlog.ProblemStatus,
				Difficulty: auxiliaryBlog.ProblemDifficulty,
			}
		}
		blogs = append(blogs, blog)
	}

	return blogs, nil
}

// 根据ID更新博客
func UpdateBlogById(b entity.Blog) error {
	tx := db.Db.Model(&b).Where("id = ?", b.Id).Updates(b)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除博客
func DeleteBlogById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Blog{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计博客数量
func CountBlogs(condition model.BlogWhere) (uint64, error) {
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

// 根据创建时间统计博客数量
func CountBlogsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var counts []model.CountByDate

	tx := db.Db.Model(&entity.Blog{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

func blogUnionJoins(tx *gorm.DB) *gorm.DB {
	query := []string{"tbl_blog.*"}
	query = append(query, briefUserSelect()...)
	query = append(query, briefProblemSelect()...)
	tx = tx.Select(query)
	tx = briefProblemJoins(tx, "tbl_blog")
	tx = briefUserJoins(tx, "tbl_blog")
	return tx
}

func briefBlogSelect() []string {
	return []string{
		"tbl_blog.title as blog_title",
		"tbl_blog.status as blog_status",
	}
}

func briefBlogJoins(db *gorm.DB, tbl string) *gorm.DB {
	return db.Joins(fmt.Sprintf("LEFT JOIN tbl_blog ON %s.blog_id = tbl_blog.id", tbl))
}
