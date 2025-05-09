package dto

import (
	"STUOJ/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 时间范围
type Period struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// 检查时间范围
func (p *Period) Check() error {
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

	return nil
}

func (p *Period) GetPeriod(c *gin.Context) error {
	var err error

	// 读取参数
	startTimeStr := c.Query("start-time")
	endTimeStr := c.Query("end-time")

	// 解析时间范围
	err = p.FromString(startTimeStr, endTimeStr, utils.DATETIME_LAYOUT)
	if err != nil {
		return err
	}

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return err
	}

	return nil
}
