package entity

import "time"

// LanguageStatus 语言状态：1 弃用，2 禁用，3 启用
//
//go:generate go run ../../../../dev/gen/enum_valid.go
type LanguageStatus uint8

const (
	LanguageDeprecated LanguageStatus = 1
	LanguageDisabled   LanguageStatus = 2
	LanguageEnabled    LanguageStatus = 3
)

func (s LanguageStatus) String() string {
	switch s {
	case LanguageDeprecated:
		return "弃用"
	case LanguageDisabled:
		return "禁用"
	case LanguageEnabled:
		return "启用"
	default:
		return "未知"
	}
}

// Language 编程语言
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Language
//go:generate go run ../../../../dev/gen/field_select.go -struct=Language
type Language struct {
	Id         uint64         `gorm:"primaryKey;autoIncrement;comment:语言Id"`
	Name       string         `gorm:"type:varchar(255);not null;comment:语言名"`
	Serial     uint16         `gorm:"not null;default:0;comment:排序序号"`
	MapId      uint32         `gorm:"not null;default:0;comment:映射Id"`
	Status     LanguageStatus `gorm:"not null;default:3;comment:状态"`
	CreateTime time.Time      `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time      `gorm:"autoUpdateTime;comment:更新时间"`
}

func (Language) TableName() string {
	return "tbl_language"
}
