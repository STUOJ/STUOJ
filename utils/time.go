package utils

import (
	"time"
)

const (
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
	DATE_LAYOUT     = "2006-01-02"
)

func GenerateDateList(startDate time.Time, endDate time.Time) []string {
	var dateList []string
	for startDate.Before(endDate.AddDate(0, 0, 1)) {
		dateList = append(dateList, startDate.Format(DATE_LAYOUT))
		startDate = startDate.AddDate(0, 0, 1)
	}
	return dateList
}
