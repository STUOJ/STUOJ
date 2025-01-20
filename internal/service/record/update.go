package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 更新提交记录
func UpdateSubmission(s entity.Submission) error {
	updateTime := time.Now()
	s.UpdateTime = updateTime

	err := dao.UpdateSubmissionById(s)
	if err != nil {
		return err
	}

	return nil
}

// 更新评测结果
func UpdateJudgement(j entity.Judgement) error {
	err := dao.UpdateJudgementById(j)
	if err != nil {
		return err
	}

	// 更新提交记录状态更新时间
	err = dao.UpdateSubmissionUpdateTimeById(j.SubmissionId)
	if err != nil {
		log.Println(err)
		return errors.New("更新提交记录状态更新时间失败")
	}

	return nil
}
