package collection

/*
func GetStatistics() (model.CollectionStatistics, error) {
	var err error
	var stats model.CollectionStatistics
	var cbps []model.CountByCollection

	// 统计题单数量
	stats.CollectionCount, err = dao.CountCollections(model.CollectionWhere{})
	if err != nil {
		log.Println(err)
		return model.CollectionStatistics{}, errors.New("统计题单数量失败")
	}

	// 统计题目数量
	cbps, err = dao.CountProblemsGroupByCollection()
	if err != nil {
		log.Println(err)
		return model.CollectionStatistics{}, errors.New("统计题目数量失败")
	}
	stats.ProblemCountByCollection = mapCountFromCountByCollection(cbps)

	return stats, nil
}

func mapCountFromCountByCollection(cbts []model.CountByCollection) model.MapCount {
	m := make(model.MapCount)
	for _, v := range cbts {
		var collection entity.Collection
		collection, err := dao.SelectBlogById(v.CollectionId)
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
*/
