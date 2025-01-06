package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

type SubmissionPage struct {
	Submissions []entity.Submission `json:"submissions"`
	model.Page
}

// 查询所有提交记录（不返回源代码）
func Select(condition dao.SubmissionWhere, userId uint64, hideCode ...bool) (SubmissionPage, error) {
	if !condition.Page.Exist() {
		condition.Page.Set(1)
	}
	if !condition.Size.Exist() {
		condition.Size.Set(10)
	}
	// 获取提交信息
	submissions, err := dao.SelectSubmissions(condition)
	if err != nil {
		log.Println(err)
		return SubmissionPage{}, errors.New("获取提交信息失败")
	}
	if len(hideCode) == 0 || hideCode[0] { // 隐藏源代码
		hideSubmissionSourceCode(userId, submissions)
	}
	total, err := dao.CountSubmissions(condition)
	if err != nil {
		log.Println(err)
		return SubmissionPage{}, errors.New("获取提交记录总数失败")
	}
	sPage := SubmissionPage{
		Submissions: submissions,
		Page: model.Page{
			Page:  condition.Page.Value(),
			Size:  condition.Size.Value(),
			Total: total,
		},
	}
	return sPage, nil
}

// 根据提交ID查询提交记录
func SelectBySubmissionId(userId uint64, sid uint64, hideCode ...bool) (model.Record, error) {
	// 获取提交信息
	s, err := dao.SelectSubmissionById(sid)
	if err != nil {
		log.Println(err)
		return model.Record{}, errors.New("获取提交信息失败")
	}

	// 获取评测结果
	judgements, err := dao.SelectJudgementsBySubmissionId(sid)
	if err != nil {
		log.Println(err)
		return model.Record{}, errors.New("获取评测结果失败")
	}

	if (len(hideCode) == 0 || hideCode[0]) && userId != s.UserId { // 隐藏源代码
		s.SourceCode = ""
	}

	// 封装提交记录
	r := model.Record{
		Submission: s,
		Judgements: judgements,
	}

	return r, nil
}

// 隐藏源代码
func hideSubmissionSourceCode(userId uint64, submissions []entity.Submission) {
	for i := range submissions {
		if submissions[i].UserId != userId {
			submissions[i].SourceCode = ""
		}
	}
}

// 封装提交记录
// func wrapRecords(submissions []entity.Submission) []model.Record {
// 	var records []model.Record

// 	hideSubmissionSourceCode(submissions)

// 	for _, s := range submissions {
// 		r := model.Record{
// 			Submission: s,
// 		}
// 		records = append(records, r)
// 	}

// 	return records
// }
