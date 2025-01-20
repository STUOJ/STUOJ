package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BriefProblem struct {
	ProblemTitle      string               `gorm:"column:problem_title"`
	ProblemStatus     entity.ProblemStatus `gorm:"column:problem_status"`
	ProblemDifficulty entity.Difficulty    `gorm:"column:problem_difficulty"`
}

type auxiliaryProblem struct {
	entity.Problem
	ProblemHistoryUserId string `gorm:"column:problem_history_user_id"`
	ProblemTagIds        string `gorm:"column:problem_tag_id"`
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
	var p auxiliaryProblem

	subQueryUser := db.Db.Model(&entity.History{}).
		Select("GROUP_CONCAT(DISTINCT user_id)").
		Where("problem_id = ?", id)

	subQueryTag := db.Db.Model(&entity.ProblemTag{}).
		Select("GROUP_CONCAT(DISTINCT tag_id)").
		Where("problem_id = ?", id)

	tx := db.Db.Model(&entity.Problem{}).
		Where("id = ?", id).
		Select("tbl_problem.*, (?) as problem_history_user_id, (?)  as problem_tag_id", subQueryUser, subQueryTag).
		Scan(&p)

	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	// 将逗号分隔的字符串转换为 []uint64
	historyUserIds := make([]uint64, 0)
	if p.ProblemHistoryUserId != "" {
		for _, idStr := range strings.Split(p.ProblemHistoryUserId, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				historyUserIds = append(historyUserIds, id)
			}
		}
	}
	p.Problem.UserIds = historyUserIds

	tagIds := make([]uint64, 0)
	if p.ProblemTagIds != "" {
		for _, idStr := range strings.Split(p.ProblemTagIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				tagIds = append(tagIds, id)
			}
		}
	}
	p.Problem.TagIds = tagIds

	return p.Problem, nil
}

func SelectProblems(condition model.ProblemWhere) ([]entity.Problem, error) {
	var problems []auxiliaryProblem

	subQueryTag := db.Db.Model(&entity.ProblemTag{}).
		Select("GROUP_CONCAT(DISTINCT tag_id)").
		Where("problem_id = tbl_problem.id")

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.
		Select("tbl_problem.*, (?) as problem_tag_id", subQueryTag).
		Scan(&problems)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// 处理每个问题的标签
	for i := range problems {
		tagIds := make([]uint64, 0)
		if problems[i].ProblemTagIds != "" {
			for _, idStr := range strings.Split(problems[i].ProblemTagIds, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					tagIds = append(tagIds, id)
				}
			}
		}
		problems[i].Problem.TagIds = tagIds
	}

	// 将辅助结构体转换为实体结构体
	result := make([]entity.Problem, len(problems))
	for i := range problems {
		result[i] = problems[i].Problem
	}

	return result, nil
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
func CountProblems(condition model.ProblemWhere) (uint64, error) {
	var count int64

	where := condition.GenerateWhereWithNoPage()

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

func briefProblemSelect() []string {
	return []string{
		"tbl_problem.title as problem_title",
		"tbl_problem.status as problem_status",
		"tbl_problem.difficulty as problem_difficulty",
	}
}

func briefProblemJoins(db *gorm.DB, tbl string) *gorm.DB {
	db = db.Joins(fmt.Sprintf("LEFT JOIN tbl_problem ON %s.problem_id = tbl_problem.id", tbl))
	return db
}
