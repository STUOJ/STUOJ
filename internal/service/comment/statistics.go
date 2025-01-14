package comment

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 获取评论数量
func GetStatisticsOfSubmitByPeriod(p model.Period) (model.MapCount, error) {
	var err error

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计博客数量
	cbds, err := dao.CountCommentsBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计评论数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}
