package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type BlogWhere struct {
	Id        model.Field[uint64]
	UserId    model.Field[uint64]
	ProblemId model.Field[uint64]
	Title     model.Field[string]
	Status    model.Field[entity.BlogStatus]
	StartTime model.Field[time.Time]
	EndTime   model.Field[time.Time]
	Page      model.Field[uint64]
	Size      model.Field[uint64]
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
	var b entity.Blog

	tx := db.Db.Where("id = ?", id).First(&b)
	if tx.Error != nil {
		return entity.Blog{}, tx.Error
	}

	return b, nil
}

func SelectBlogs(condition BlogWhere) ([]entity.Blog, error) {
	var blogs []entity.Blog

	where := generateBlogWhereCondition(condition)

	tx := db.Db.Model(&entity.Blog{})
	tx = where(tx)
	tx = tx.Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
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
func CountBlogs(condition BlogWhere) (uint64, error) {
	var count int64

	where := generateBlogWhereConditionWithNoPage(condition)
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

func generateBlogWhereConditionWithNoPage(con BlogWhere) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}
		if con.Id.Exist() {
			whereClause["id"] = con.Id.Value()
		}
		if con.ProblemId.Exist() {
			whereClause["problem_id"] = con.ProblemId.Value()
		}
		if con.UserId.Exist() {
			whereClause["user_id"] = con.UserId.Value()
		}
		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		where := db.Where(whereClause)

		if con.Title.Exist() {
			where = where.Where("title LIKE ?", "%"+con.Title.Value()+"%")
		}
		if con.StartTime.Exist() {
			where.Where("create_time >= ?", con.StartTime.Value())
		}
		if con.EndTime.Exist() {
			where.Where("create_time <= ?", con.EndTime.Value())
		}
		return where
	}
}

func generateBlogWhereCondition(con BlogWhere) func(*gorm.DB) *gorm.DB {
	where := generateBlogWhereConditionWithNoPage(con)
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
