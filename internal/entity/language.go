package entity

// Language 编程语言
type Language struct {
	Id     uint64 `gorm:"primaryKey;autoIncrement;comment:语言ID" json:"id"`
	Name   string `gorm:"type:varchar(255);not null;comment:语言名" json:"name"`
	Serial uint64 `gorm:"not null;default:0;comment:排序序号" json:"serial"`
	MapId  int    `gorm:"not null;default:0;comment:映射ID" json:"map_id"`
	Status uint64 `gorm:"not null;default:3;comment:状态" json:"status"`
}

func (Language) TableName() string {
	return "tbl_language"
}
