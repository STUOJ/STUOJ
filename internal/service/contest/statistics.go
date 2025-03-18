package contest

/*
func GetStatistics() (model.ContestStatistics, error) {
	var err error
	var stats model.ContestStatistics
	var cbps []model.CountByContest

	// 统计比赛数量
	stats.ContestCount, err = dao.CountContests(model.ContestWhere{})
	if err != nil {
		log.Println(err)
		return model.ContestStatistics{}, errors.New("统计比赛数量失败")
	}

	// 统计题目数量
	cbps, err = dao.CountProblemsGroupByContest()
	if err != nil {
		log.Println(err)
		return model.ContestStatistics{}, errors.New("统计题目数量失败")
	}
	stats.ProblemCountByContest = mapCountFromCountByContest(cbps)

	return stats, nil
}

func mapCountFromCountByContest(cbts []model.CountByContest) model.MapCount {
	m := make(model.MapCount)
	for _, v := range cbts {
		var contest entity.Contest
		contest, err := dao.SelectBlogById(v.ContestId)
		if err != nil {
			log.Println(err)
			contest = entity.Contest{
				Id:   v.ContestId,
				Name: "未知比赛",
			}
		}
		m[contest.Name] = v.Count
	}

	return m
}
*/
