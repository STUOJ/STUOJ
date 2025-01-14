package model

import (
	"STUOJ/internal/entity"
	"STUOJ/utils"
	"time"
)

type MapCount map[string]int64

const LayoutCountByDate = "2006-01-02"

// 按日期统计
type CountByDate struct {
	Date  time.Time `json:"date"`
	Count int64     `json:"count"`
}

func (m *MapCount) FromCountByDate(cbds []CountByDate) {
	*m = make(MapCount)
	for _, v := range cbds {
		date := v.Date.Format(LayoutCountByDate)
		(*m)[date] = v.Count
	}
}

// 按角色统计
type CountByRole struct {
	Role  entity.Role `json:"role"`
	Count int64       `json:"count"`
}

func (m *MapCount) FromCountByRole(cbrs []CountByRole) {
	*m = make(MapCount)
	for _, v := range cbrs {
		(*m)[v.Role.String()] = v.Count
	}
}

// 按评测状态统计
type CountByJudgeStatus struct {
	Status entity.JudgeStatus `json:"status"`
	Count  int64              `json:"count"`
}

func (m *MapCount) FromCountByJudgeStatus(cbjss []CountByJudgeStatus) {
	*m = make(MapCount)
	for _, v := range cbjss {
		(*m)[v.Status.String()] = v.Count
	}
}

// 按标签统计
type CountByTag struct {
	TagId uint64 `json:"tag_id"`
	Count int64  `json:"count"`
}

// 按语言统计
type CountByLanguage struct {
	LanguageId uint64 `json:"language_id"`
	Count      int64  `json:"count"`
}

func (m *MapCount) MapCountFillZero(startDate time.Time, endDate time.Time) {
	dateList := utils.GenerateDateList(startDate, endDate)
	// 填充没有结果的日期
	for _, date := range dateList {
		if _, ok := (*m)[date]; !ok {
			(*m)[date] = 0
		}
	}
}
