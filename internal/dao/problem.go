package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type ProblemWhere struct {
	Id         model.Field[uint64]
	Title      model.Field[string]
	Difficulty model.Field[entity.Difficulty]
	Status     model.Field[entity.ProblemStatus]
	Tag        model.FieldList[uint64]
	Page       model.Field[uint64]
	Size       model.Field[uint64]
}

// 插入题目
func InsertProblem(p entity.Problem) (uint64, error) {
	tx := db.Db.Create(&p)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return p.Id, nil
}

func SelectProblemById(id uint64) (entity.Problem, error) {
	var p entity.Problem

	tx := db.Db.Where("id = ?", id).First(&p)
	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	return p, nil
}

func SelectProblem(condition ProblemWhere) ([]entity.Problem, error) {
	var problems []entity.Problem
	where := generateProblemWhereCondition(condition)
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.Find(&problems)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return problems, nil
}

// 根据ID更新题目
func UpdateProblemById(p entity.Problem) error {
	tx := db.Db.Model(&p).Where("id = ?", p.Id).Updates(p)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID更新题目更新时间
func UpdateProblemUpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Problem{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除题目
func DeleteProblemById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Problem{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计题目数量
func CountProblems(condition ProblemWhere) (uint64, error) {
	var count int64

	where := generateProblemWhereConditionWithNoPage(condition)

	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 根据创建时间统计用户数量
func CountProblemsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.Problem{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}

func generateProblemWhereConditionWithNoPage(con ProblemWhere) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		whereClause := map[string]interface{}{}

		if con.Id.Exist() {
			whereClause["id"] = con.Id.Value()
		}
		if con.Status.Exist() {
			whereClause["status"] = con.Status.Value()
		}
		if con.Difficulty.Exist() {
			whereClause["difficulty"] = con.Difficulty.Value()
		}

		where := db.Where(whereClause)
		if con.Tag.Exist() {
			where = where.Where("id IN (SELECT problem_id FROM tbl_problem_tag WHERE tag_id In(?) GROUP BY problem_id HAVING COUNT(DISTINCT tag_id) =?)", con.Tag.Value(), len(con.Tag.Value()))
		}
		if con.Title.Exist() {
			where = where.Where("title LIKE ?", "%"+con.Title.Value()+"%")
		}
		return where
	}
}

func generateProblemWhereCondition(con ProblemWhere) func(*gorm.DB) *gorm.DB {
	where := generateProblemWhereConditionWithNoPage(con)
	return func(db *gorm.DB) *gorm.DB {
		return where(db).Offset(int((con.Page.Value() - 1) * con.Size.Value())).Limit(int(con.Size.Value()))
	}
}
