package entity

import "time"

// BlogStatus 状态：1 屏蔽, 2 草稿, 3 公开, 4 公告
//go:generate go run ../../../utils/gen/enum_valid.go BlogStatus
type BlogStatus uint8

const (
	BlogBanned BlogStatus = 1
	BlogDraft  BlogStatus = 2
	BlogPublic BlogStatus = 3
	BlogNotice BlogStatus = 4
)

func (s BlogStatus) String() string {
	switch s {
	case BlogBanned:
		return "屏蔽"
	case BlogPublic:
		return "公开"
	case BlogDraft:
		return "草稿"
	case BlogNotice:
		return "公告"
	default:
		return "未知"
	}
}

//go:generate go run ../../../utils/gen/dao_store.go -struct=Blog
//go:generate go run ../../../utils/gen/field_select.go -struct=Blog
// Blog 博客
type Blog struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement;comment:博客ID"`
	UserId     uint64     `gorm:"not null;default:0;comment:用户ID"`
	ProblemId  uint64     `gorm:"not null;default:0;comment:关联题目ID"`
	Title      string     `gorm:"type:text;not null;comment:标题"`
	Content    string     `gorm:"type:longtext;not null;comment:内容"`
	Status     BlogStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	User       User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
