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

// 按语言统计提交记录数量
func GetStatisticsOfSubmissionLanguage() (model.MapCount, error) {
	var err error

	// 按语言统计提交记录数量
	cbls, err := dao.CountSubmissionsGroupByLanguageId()
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计提交记录失败")
	}

	mc := mapCountFromCountByLanguage(cbls)

	return mc, nil
}

// 按评测状态统计提交记录数量
func GetStatisticsOfSubmissionStatus() (model.MapCount, error) {
	var err error

	// 按评测状态统计提交记录数量
	cbjss, err := dao.CountSubmissionsGroupByStatus()
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计提交记录失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByJudgeStatus(cbjss)

	return mc, nil
}

// 按评测状态统计评测结果数量
func GetStatisticsOfJudgementStatus() (model.MapCount, error) {
	var err error

	// 按评测状态统计评测结果数量
	cbjss, err := dao.CountJudgementsGroupByStatus()
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计评测结果失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByJudgeStatus(cbjss)

	return mc, nil
}

// 提交记录统计
func GetStatisticsOfSubmitByPeriod(p model.Period) (model.MapCount, error) {
	var err error

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 按日期统计提交记录数量
	cbds, err := dao.CountSubmissionsBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计提交记录失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
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
