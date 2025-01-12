package model

import (
	"errors"
	"log"
	"time"
)

// 时间范围
type Period struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// 检查时间范围
func (p Period) Check() error {
	if p.StartTime.After(p.EndTime) {
		return errors.New("开始时间不能晚于结束时间")
	}
	return nil
}

// 从字符串解析时间范围
func (p *Period) FromString(startTimeStr string, endTimeStr string, layout string) error {
	var err error

	if startTimeStr == "" || endTimeStr == "" {
		return errors.New("参数错误")
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return errors.New("加载时区失败")
	}

	p.StartTime, err = time.ParseInLocation(layout, startTimeStr, location)
	if err != nil {
		log.Println(err)
		return errors.New("开始时间格式错误")
	}
	p.EndTime, err = time.ParseInLocation(layout, endTimeStr, location)
	if err != nil {
		log.Println(err)
		return errors.New("结束时间格式错误")
	}

	log.Println(p.StartTime, p.EndTime)

	return nil
}
