package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"errors"
	"log"
)

// 统计用户
func GetStatistics() (model.UserStatistics, error) {
	var err error
	var stats model.UserStatistics

	// 统计用户数量
	stats.UserCount, err = dao.CountUsers(model.UserWhere{})
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户数量失败")
	}

	return stats, nil
}

// 统计用户
func GetStatisticsOfRole() (model.MapCount, error) {
	var err error

	// 统计用户角色
	cbrs, err := dao.CountUsersGroupByRole()
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计用户角色失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByRole(cbrs)

	return mc, nil
}

// 统计用户
func GetStatisticsOfRegisterByPeriod(p model.Period) (model.MapCount, error) {
	var err error

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.MapCount{}, err
	}

	// 统计用户注册数量
	cbds, err := dao.CountUsersBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.MapCount{}, errors.New("统计用户注册数量失败")
	}

	mc := make(model.MapCount)
	mc.FromCountByDate(cbds)
	mc.MapCountFillZero(p.StartTime, p.EndTime)

	return mc, nil
}
