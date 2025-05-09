package valueobject

import (
	"STUOJ/internal/model"
	"errors"
	"time"
)

type EndTime struct {
	model.Valueobject[time.Time]
}

func NewEndTime(value time.Time) EndTime {
	var endTime EndTime
	endTime.Set(value)
	return endTime
}

func (endTime EndTime) Verify() error {
	if endTime.Value().Before(time.Now()) {
		return errors.New("比赛结束时间必须晚于当前时间！")
	}
	return nil
}
