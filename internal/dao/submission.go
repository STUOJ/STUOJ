package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type auxiliarySubmission struct {
	entity.Submission
	model.BriefUser
	model.BriefProblem
}

// 插入提交记录
func InsertSubmission(s entity.Submission) (uint64, error) {
	updateTime := time.Now()
	s.UpdateTime = updateTime
	s.CreateTime = updateTime
	tx := db.Db.Create(&s)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return s.Id, nil
}

// 查询提交记录
func SelectSubmissions(condition model.SubmissionWhere) ([]entity.Submission, error) {
	var auxiliarySubmissions []auxiliarySubmission
	var submissions []entity.Submission

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Submission{})
	tx = where(tx)
	tx = tx.Find(&auxiliarySubmissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, auxiliarySubmission := range auxiliarySubmissions {
		submission := auxiliarySubmission.Submission
		submission.User = entity.User{
			Id:       auxiliarySubmission.UserId,
			Username: auxiliarySubmission.Username,
			Role:     auxiliarySubmission.Role,
			Avatar:   auxiliarySubmission.Avatar,
		}
		submission.Problem = entity.Problem{
			Id:         auxiliarySubmission.ProblemId,
			Title:      auxiliarySubmission.ProblemTitle,
			Status:     auxiliarySubmission.ProblemStatus,
			Difficulty: auxiliarySubmission.ProblemDifficulty,
		}
		submissions = append(submissions, submission)
	}

	return submissions, nil
}

// 根据ID查询提交记录
func SelectSubmissionById(id uint64) (entity.Submission, error) {
	var auxiliarySubmission auxiliarySubmission
	var s entity.Submission

	condition := model.SubmissionWhere{}

	where := condition.GenerateWhere()

	tx := db.Db.Where(&entity.Submission{Id: id})
	tx = where(tx)
	tx = tx.First(&auxiliarySubmission)
	if tx.Error != nil {
		return entity.Submission{}, tx.Error
	}
	s = auxiliarySubmission.Submission
	s.User = entity.User{
		Id:       auxiliarySubmission.UserId,
		Username: auxiliarySubmission.Username,
		Role:     auxiliarySubmission.Role,
		Avatar:   auxiliarySubmission.Avatar,
	}
	s.Problem = entity.Problem{
		Id:         auxiliarySubmission.ProblemId,
		Title:      auxiliarySubmission.ProblemTitle,
		Status:     auxiliarySubmission.ProblemStatus,
		Difficulty: auxiliarySubmission.ProblemDifficulty,
	}

	return s, nil
}

// 更新提交记录
func UpdateSubmissionById(s entity.Submission) error {
	tx := db.Db.Model(&s).Updates(s)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func UpdateSubmissionByIdAndInsertJudgements(s entity.Submission, j []entity.Judgement) error {
	tx := db.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&s).Updates(s).Error; err != nil {
			return err
		}
		for _, j := range j {
			if err := tx.Create(&j).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return tx
}

// 根据ID更新提交记录的更新时间
func UpdateSubmissionUpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Submission{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除提交记录
func DeleteSubmissionById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Submission{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计提交信息数量
func CountSubmissions(condition model.SubmissionWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.Submission{})
	tx = where(tx)
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}

// 按评测状态统计提交信息数量
func CountSubmissionsGroupByStatus() ([]model.CountByJudgeStatus, error) {
	var counts []model.CountByJudgeStatus

	tx := db.Db.Model(&entity.Submission{}).Select("status, count(*) as count").Group("status").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

// 按语言ID统计提交信息数量
func CountSubmissionsGroupByLanguageId() ([]model.CountByLanguage, error) {
	var counts []model.CountByLanguage

	tx := db.Db.Model(&entity.Submission{}).Select("language_id, count(*) as count").Group("language_id").Scan(&counts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return counts, nil
}

// 根据创建时间统计用户数量
func CountSubmissionsBetweenCreateTime(startTime time.Time, endTime time.Time) ([]model.CountByDate, error) {
	var countByDate []model.CountByDate

	tx := db.Db.Model(&entity.Submission{}).Where("create_time between ? and ?", startTime, endTime).Select("date(create_time) as date, count(*) as count").Group("date").Scan(&countByDate)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return countByDate, nil
}
