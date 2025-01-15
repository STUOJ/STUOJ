package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

func GetStatistics() (model.TagStatistics, error) {
	var err error
	var stats model.TagStatistics
	var cbts []model.CountByTag

	// 统计标签数量
	stats.TagCount, err = dao.CountTags(model.TagWhere{})
	if err != nil {
		log.Println(err)
		return model.TagStatistics{}, errors.New("统计标签数量失败")
	}

	// 统计题目数量
	cbts, err = dao.CountProblemsGroupByTag()
	if err != nil {
		log.Println(err)
		return model.TagStatistics{}, errors.New("统计题目数量失败")
	}
	stats.ProblemCountByTag = mapCountFromCountByTag(cbts)

	return stats, nil
}

func mapCountFromCountByTag(cbts []model.CountByTag) model.MapCount {
	m := make(model.MapCount)
	for _, v := range cbts {
		var tag entity.Tag
		tag, err := dao.SelectTagById(v.TagId)
		if err != nil {
			log.Println(err)
			tag = entity.Tag{
				Id:   v.TagId,
				Name: "未知标签",
			}
		}
		m[tag.Name] = v.Count
	}

	return m
}
