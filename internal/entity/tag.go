package entity

// 标签
type Tag struct {
	Id       uint64    `gorm:"primaryKey;autoIncrement;comment:标签ID" json:"id,omitempty"`
	Name     string    `gorm:"type:varchar(255);not null;unique;default:'';comment:标签名" json:"name,omitempty"`
	Problems []Problem `gorm:"-" json:"-"`
}

func (Tag) TableName() string {
	return "tbl_tag"
}
