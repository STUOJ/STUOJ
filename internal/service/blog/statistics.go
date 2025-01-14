package blog

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 统计博客数量
func GetStatistics(conditon model.BlogWhere) (model.BlogStatistics, error) {
	var err error
	var stats model.BlogStatistics

	// 统计博客数量
	stats.BlogCount, err = dao.CountBlogs(conditon)
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计博客数量失败")
	}

	// 统计评论数量
	stats.CommentCount, err = dao.CountBlogs(conditon)
	if err != nil {
		log.Println(err)
		return model.BlogStatistics{}, errors.New("统计评论数量失败")
	}

	return stats, nil
}

// 统计发表博客数量
func GetStatisticsOfSubmitByPeriod(p model.Period) (model.MapCount, error) {
	var err error

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计博客数量
	cbds, err := dao.CountBlogsBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计博客数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}
