package entity

// LanguageStatus 语言状态：1 弃用，2 禁用，3 启用
//go:generate go run ../../utils/gen/enum_valid.go LanguageStatus entity
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
type Language struct {
	Id     uint64         `gorm:"primaryKey;autoIncrement;comment:语言ID"`
	Name   string         `gorm:"type:varchar(255);not null;comment:语言名"`
	Serial uint16         `gorm:"not null;default:0;comment:排序序号"`
	MapId  uint32         `gorm:"not null;default:0;comment:映射ID"`
	Status LanguageStatus `gorm:"not null;default:3;comment:状态"`
}

func (Language) TableName() string {
	return "tbl_language"
}
