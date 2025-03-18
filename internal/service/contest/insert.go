package contest

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"errors"
	"log"
	"time"
)

// 插入比赛
func Insert(ct entity.Contest) (uint64, error) {
	var err error

	updateTime := time.Now()
	ct.UpdateTime = updateTime
	ct.CreateTime = updateTime

	// 插入比赛
	ct.Id, err = dao.InsertContest(ct)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入失败")
	}

	return ct.Id, nil
}
