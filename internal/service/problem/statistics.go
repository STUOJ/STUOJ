package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 统计题目
func GetStatistics() (model.ProblemStatistics, error) {
	var err error
	var stats model.ProblemStatistics

	// 统计题目数量
	stats.ProblemCount, err = dao.CountProblems(model.ProblemWhere{})
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计题目数量失败")
	}

	// 统计评测点数量
	stats.TestcaseCount, err = dao.CountTestcases()
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计评测点数量失败")
	}

	// 统计题解数量
	stats.SolutionCount, err = dao.CountBlogs(model.BlogWhere{})
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计题解数量失败")
	}

	return stats, nil
}

// 统计添加题目数量
func GetStatisticsOfInsertByPeriod(p model.Period) (model.MapCount, error) {
	var err error
	var cbds []model.CountByDate

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计添加题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationInsert)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计添加题目数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}

// 统计更新题目数量
func GetStatisticsOfUpdateByPeriod(p model.Period) (model.MapCount, error) {
	var err error
	var cbds []model.CountByDate

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计更新题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationUpdate)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计更新题目数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}

// 统计删除题目数量
func GetStatisticsOfDeleteByPeriod(p model.Period) (model.MapCount, error) {
	var err error
	var cbds []model.CountByDate

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计删除题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationDelete)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计删除题目数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}
