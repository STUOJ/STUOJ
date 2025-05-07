package entity

import "time"

// BlogStatus 状态：1 删除, 2 草稿, 3 公开, 4 公告
//
//go:generate go run ../../../../dev/gen/enum_valid.go BlogStatus
type BlogStatus uint8

const (
	BlogDeleted BlogStatus = 1
	BlogDraft   BlogStatus = 2
	BlogPublic  BlogStatus = 3
	BlogNotice  BlogStatus = 4
)

func (s BlogStatus) String() string {
	switch s {
	case BlogDeleted:
		return "删除"
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

// Blog 博客
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Blog
//go:generate go run ../../../../dev/gen/field_select.go -struct=Blog
type Blog struct {
	Id         uint64     `gorm:"primaryKey;autoIncrement;comment:博客Id"`
	UserId     uint64     `gorm:"not null;default:0;comment:用户Id"`
	ProblemId  uint64     `gorm:"not null;default:0;comment:关联题目Id"`
	Title      string     `gorm:"type:text;not null;comment:标题"`
	Content    string     `gorm:"type:longtext;not null;comment:内容"`
	Status     BlogStatus `gorm:"not null;default:1;comment:状态"`
	CreateTime time.Time  `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime time.Time  `gorm:"autoUpdateTime;comment:更新时间"`
	User       User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
}

func (Blog) TableName() string {
	return "tbl_blog"
}
