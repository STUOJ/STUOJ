package entity

// 标签
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=Tag
//go:generate go run ../../../utils/gen/field_select.go -struct=Tag
type Tag struct {
	Id   uint64 `gorm:"primaryKey;autoIncrement;comment:标签Id"`
	Name string `gorm:"type:varchar(255);not null;unique;default:'';comment:标签名"`
}

func (Tag) TableName() string {
	return "tbl_tag"
}
