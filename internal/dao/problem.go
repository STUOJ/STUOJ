package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"strconv"
	"strings"
	"time"
)

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

func SelectProblemById(id uint64, condition model.ProblemWhere) (entity.Problem, error) {
	var p auxiliaryProblem

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
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

	if condition.Testcases.Exist() && condition.Testcases.Value() {
		p.Problem.Testcases = make([]entity.Testcase, 0)
		tx := db.Db.Model(&entity.Testcase{}).Where("problem_id = ?", id).Find(&p.Problem.Testcases)
		if tx.Error != nil {
			return entity.Problem{}, tx.Error
		}
	}

	if condition.Solutions.Exist() && condition.Solutions.Value() {
		p.Problem.Solutions = make([]entity.Solution, 0)
		tx := db.Db.Model(&entity.Solution{}).Where("problem_id = ?", id).Find(&p.Problem.Solutions)
		if tx.Error != nil {
			return entity.Problem{}, tx.Error
		}
	}

	p.Problem.UserScore = p.ProblemUserScore
	p.Problem.HasUserSubmission = p.HasUserSubmission

	return p.Problem, nil
}

func SelectProblems(condition model.ProblemWhere) ([]entity.Problem, error) {
	var problems []auxiliaryProblem

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Problem{})
	tx = where(tx)
	tx = tx.Scan(&problems)

	if tx.Error != nil {
		return nil, tx.Error
	}

	// 处理每个问题的标签
	for i := range problems {
		problems[i].Problem.TagIds = make([]uint64, 0)
		if problems[i].ProblemTagIds != "" {
			for _, idStr := range strings.Split(problems[i].ProblemTagIds, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					problems[i].Problem.TagIds = append(problems[i].Problem.TagIds, id)
				}
			}
		}
		problems[i].Problem.UserIds = make([]uint64, 0)
		if problems[i].ProblemUserId != "" {
			for _, idStr := range strings.Split(problems[i].ProblemUserId, ",") {
				if id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64); err == nil {
					problems[i].Problem.UserIds = append(problems[i].Problem.UserIds, id)
				}
			}
		}
		problems[i].Problem.UserScore = problems[i].ProblemUserScore
		problems[i].Problem.HasUserSubmission = problems[i].HasUserSubmission
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
	// 明确指定要更新的字段，包含需要处理空值的字段
	tx := db.Db.Model(&p).
		Select("title", "source", "difficulty", "time_limit", "memory_limit", "description", "input", "output", "sample_input", "sample_output", "hint", "status", "update_time").
		Updates(p)
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
