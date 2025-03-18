package contest

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入比赛
func Insert(c entity.Contest) (uint64, error) {
	var err error

	updateTime := time.Now()
	c.UpdateTime = updateTime
	c.CreateTime = updateTime

	// 插入比赛
	c.Id, err = dao.InsertContest(c)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败")
	}

	return c.Id, nil
}
