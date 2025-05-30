package valueobject

import (
	"STUOJ/internal/domain/shared"
	"errors"
	"time"
)

type StartTime struct {
	shared.Valueobject[time.Time]
}

func NewStartTime(value time.Time) StartTime {
	var startTime StartTime
	startTime.Set(value)
	return startTime
}

func (startTime StartTime) Verify() error {
	if startTime.Value().Before(time.Now()) {
		return errors.New("比赛开始时间不能早于当前时间！")
	}
	return nil
}
