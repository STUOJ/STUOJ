package dao

import (
	"STUOJ/internal/db"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/model"
	"time"

	"gorm.io/gorm"
)

type submissionStore struct{}

var SubmissionStore = new(submissionStore)

// 插入提交记录
func (store *submissionStore) Insert(s entity.Submission) (uint64, error) {
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
func (store *submissionStore) Select(condition model.SubmissionWhere) ([]entity.Submission, error) {
	var submissions []entity.Submission

	where := condition.GenerateWhere()
	tx := db.Db.Model(&entity.Submission{})
	tx = where(tx)
	tx = tx.Find(&submissions)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return submissions, nil
}

// 根据ID查询提交记录
func (store *submissionStore) SelectById(id uint64) (entity.Submission, error) {
	var s entity.Submission

	condition := model.SubmissionWhere{}

	where := condition.GenerateWhere()

	tx := db.Db.Where(&entity.Submission{Id: id})
	tx = where(tx)
	tx = tx.First(&s)
	if tx.Error != nil {
		return entity.Submission{}, tx.Error
	}

	return s, nil
}

func (store *submissionStore) SelectACUsers(pid uint64, size uint64) ([]entity.User, error) {
	var users []entity.User
	tx := db.Db.Raw("SELECT * FROM (SELECT tbl_submission.*,tbl_user.username AS user_username,tbl_user.role AS user_role,tbl_user.avatar AS user_avatar,tbl_problem.title AS problem_title,tbl_problem.STATUS AS problem_status,tbl_problem.difficulty AS problem_difficulty,ROW_NUMBER() OVER (PARTITION BY tbl_submission.user_id ORDER BY tbl_submission.create_time ASC) AS rn FROM `tbl_submission` LEFT JOIN tbl_problem ON tbl_submission.problem_id=tbl_problem.id LEFT JOIN tbl_user ON tbl_submission.user_id=tbl_user.id WHERE `tbl_submission`.`status`=3 AND tbl_submission.problem_id = ? AND NOT EXISTS (SELECT 1 FROM tbl_history WHERE tbl_history.user_id=tbl_submission.user_id AND tbl_history.problem_id=tbl_submission.problem_id)) AS ranked_submissions WHERE rn=1 ORDER BY create_time ASC LIMIT ?;", pid, size)
	tx = tx.Scan(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

// 更新提交记录
func (store *submissionStore) UpdateById(s entity.Submission) error {
	tx := db.Db.Model(&s).Updates(s)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (store *submissionStore) UpdateByIdAndInsertJudgements(s entity.Submission, j []entity.Judgement) error {
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
func (store *submissionStore) UpdateTimeById(id uint64) error {
	tx := db.Db.Model(&entity.Submission{}).Where("id = ?", id).Update("update_time", time.Now())
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 根据ID删除提交记录
func (store *submissionStore) DeleteById(id uint64) error {
	tx := db.Db.Where("id = ?", id).Delete(&entity.Submission{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// 统计提交信息数量
func (store *submissionStore) Count(condition model.SubmissionWhere) (uint64, error) {
	var count int64
	where := condition.GenerateWhereWithNoPage()
	tx := db.Db.Model(&entity.Submission{})
	tx = where(tx)
	if condition.Distinct.Exist() {
		tx = tx.Distinct(condition.Distinct.Value())
	}
	tx = tx.Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return uint64(count), nil
}
