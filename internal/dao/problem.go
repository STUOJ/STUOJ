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
	ProblemUserId            string `gorm:"column:problem_user_id"`
	ProblemTagIds            string `gorm:"column:problem_tag_id"`
	ProblemCollectionIds     string `gorm:"column:problem_collection_id"`
	ProblemCollectionUserIds string `gorm:"column:problem_collection_user_id"`
	ProblemUserScore         uint64 `gorm:"column:problem_user_score"`
	HasUserSubmission        bool   `gorm:"column:has_user_submission"`
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

	tx := db.Db.Model(&entity.Problem{})
	tx = problemJoins(tx)
	tx = tx.Where(&entity.Problem{Id: id}).
		Scan(&p)

	if tx.Error != nil {
		return entity.Problem{}, tx.Error
	}

	// 将逗号分隔的字符串转换为 []uint64
	p.Problem.UserIds = make([]uint64, 0)
	if p.ProblemUserId != "" {
		for _, idStr := range strings.Split(p.ProblemUserId, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				p.Problem.UserIds = append(p.Problem.UserIds, id)
			}
		}
	}

	p.Problem.TagIds = make([]uint64, 0)
	if p.ProblemTagIds != "" {
		for _, idStr := range strings.Split(p.ProblemTagIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				p.Problem.TagIds = append(p.Problem.TagIds, id)
			}
		}
	}

	p.Problem.CollectionIds = make([]uint64, 0)
	if p.ProblemCollectionIds != "" {
		for _, idStr := range strings.Split(p.ProblemCollectionIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				p.Problem.CollectionIds = append(p.Problem.CollectionIds, id)
			}
		}
	}

	p.Problem.CollectionUserIds = make([]uint64, 0)
	if p.ProblemCollectionUserIds != "" {
		for _, idStr := range strings.Split(p.ProblemCollectionUserIds, ",") {
			if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
				p.Problem.CollectionUserIds = append(p.Problem.CollectionUserIds, id)
			}
		}
	}

	return p.Problem, nil
}

func SelectProblems(condition model.ProblemWhere) ([]entity.Problem, error) {
	var problems []auxiliaryProblem

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = problemJoins(tx)
	tx = tx.Scan(&problems)

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
		userIds := make([]uint64, 0)
		if problems[i].ProblemUserId != "" {
			for _, idStr := range strings.Split(problems[i].ProblemUserId, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					userIds = append(userIds, id)
				}
			}
		}
		problems[i].Problem.UserIds = userIds
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

func problemJoins(tx *gorm.DB) *gorm.DB {
	query := []string{"tbl_problem.*"}
	query = append(query,
		"(SELECT GROUP_CONCAT(DISTINCT tbl_problem_tag.tag_id) FROM tbl_problem_tag WHERE problem_id = tbl_problem.id) AS problem_tag_id",
		"(SELECT GROUP_CONCAT(DISTINCT tbl_history.user_id) FROM tbl_history WHERE problem_id = tbl_problem.id) AS problem_user_id",
		"(SELECT GROUP_CONCAT(DISTINCT tbl_collection_problem.collection_id) FROM tbl_collection_problem WHERE problem_id = tbl_problem.id) AS problem_collection_id",
		"(SELECT GROUP_CONCAT(DISTINCT tbl_collection_user.user_id) FROM tbl_collection_problem JOIN tbl_collection_user ON tbl_collection_problem.collection_id = tbl_collection_user.collection_id WHERE tbl_collection_problem.problem_id = tbl_problem.id AND EXISTS (SELECT 1 FROM tbl_collection WHERE tbl_collection.id = tbl_collection_user.collection_id AND tbl_collection.user_id IN (SELECT DISTINCT user_id FROM tbl_history WHERE problem_id = tbl_problem.id))) AS problem_collection_user_id",
	)
	tx = tx.Select(strings.Join(query, ", "))
	return tx
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
