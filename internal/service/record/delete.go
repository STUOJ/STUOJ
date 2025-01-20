package record

import (
	"STUOJ/internal/dao"
	"errors"
	"log"
)

// 根据提交ID删除提交记录
func DeleteSubmission(sid uint64) error {
	// 获取提交信息
	_, err := dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		return errors.New("提交记录不存在")
	}

	// 删除评测结果
	err = dao.DeleteJudgementBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		return errors.New("删除评测结果失败")
	}

	// 删除提交信息
	err = dao.DeleteSubmissionById(sid)
	if err != nil {
		log.Println(err)
		return errors.New("删除提交信息失败")
	}

	return nil
}
