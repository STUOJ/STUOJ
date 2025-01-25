package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"
)

type auxiliaryBlog struct {
	entity.Blog
	model.BriefUser
	model.BriefProblem
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
	condition := model.BlogWhere{}
	tx := db.Db.Model(&entity.Blog{})
	tx = tx.Where(&entity.Blog{Id: id})
	tx = condition.GenerateWhere()(tx)
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
