package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
)

// 根据邮箱验证密码
func VerifyByEmail(u entity.User) (string, error) {
	password := u.Password

	// 查询用户
	u, err := dao.SelectUserByEmail(u.Email.String())
	if err != nil {
		log.Println(err)
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = u.VerifyByPassword(password)
	if err != nil {
		log.Println(err)
		return "", errors.New("用户名或密码错误")
	}

	// 生成token
	token, err := utils.GenerateToken(u.Id)
	if err != nil {
		log.Println(err)
		return "", errors.New("生成token失败")
	}

	return token, nil
}

// 统计用户
func GetStatistics() (model.UserStatistics, error) {
	var err error
	var stats model.UserStatistics

	// 统计用户数量
	stats.UserCount, err = dao.CountUsers(dao.UserWhere{})
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
	utils.MapCountFillZero(&mc, p.StartTime, p.EndTime)

	return mc, nil
}
