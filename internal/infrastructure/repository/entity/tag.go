package entity

import "time"

// 标签
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Tag
//go:generate go run ../../../../dev/gen/field_select.go -struct=Tag
type Tag struct {
	Id         uint64    `gorm:"primaryKey;autoIncrement;comment:标签Id"`
	Name       string    `gorm:"type:varchar(255);not null;unique;default:'';comment:标签名"`
	CreateTime time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

func (Tag) TableName() string {
	return "tbl_tag"
}
