package collection

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

func GetStatistics() (model.CollectionStatistics, error) {
	var err error
	var stats model.CollectionStatistics
	var cbts []model.CountByCollection

	// 统计题单数量
	stats.CollectionCount, err = dao.CountCollections(model.CollectionWhere{})
	if err != nil {
		log.Println(err)
		return model.CollectionStatistics{}, errors.New("统计题单数量失败")
	}

	// 统计题目数量
	cbts, err = dao.CountProblemsGroupByCollection()
	if err != nil {
		log.Println(err)
		return model.CollectionStatistics{}, errors.New("统计题目数量失败")
	}
	stats.ProblemCountByCollection = mapCountFromCountByCollection(cbts)

	return stats, nil
}

func mapCountFromCountByCollection(cbts []model.CountByCollection) model.MapCount {
	m := make(model.MapCount)
	for _, v := range cbts {
		var collection entity.Collection
		collection, err := dao.SelectCollectionById(v.CollectionId)
		if err != nil {
			log.Println(err)
			collection = entity.Collection{
				Id:   v.CollectionId,
				Name: "未知题单",
			}
		}
		m[collection.Name] = v.Count
	}

	return m
}
