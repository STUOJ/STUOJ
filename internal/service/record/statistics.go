package record

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 提交记录统计
func GetStatistics(condition model.SubmissionWhere) (model.RecordStatistics, error) {
	var err error
	var stats model.RecordStatistics

	// 统计提交记录数量
	stats.SubmissionCount, err = dao.CountSubmissions(condition)
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计提交记录数量失败")
	}

	// 统计评测点结果数量
	stats.JudgementCount, err = dao.CountJudgements()
	if err != nil {
		log.Println(err)
		return model.RecordStatistics{}, errors.New("统计评测点结果数量失败")
	}

	return stats, nil
}

func mapCountFromCountByLanguage(cbts []model.CountByLanguage) model.MapCount {
	m := make(model.MapCount)
	for _, v := range cbts {
		var l entity.Language
		l, err := dao.SelectLanguageById(v.LanguageId)
		if err != nil {
			log.Println(err)
			l = entity.Language{
				Id:   v.LanguageId,
				Name: "未知语言",
			}
		}
		m[l.Name] = v.Count
	}

	return m
}
