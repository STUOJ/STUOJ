package entity

// 题目标签关系
type ProblemTag struct {
	Id        uint64  `gorm:"primaryKey;autoIncrement;comment:关系ID" json:"id"`
	ProblemId uint64  `gorm:"not null;default:0;comment:题目ID" json:"problem_id"`
	TagId     uint64  `gorm:"not null;default:0;comment:标签ID" json:"tag_id"`
	Problem   Problem `gorm:"foreignKey:ProblemID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"problem"`
	Tag       Tag     `gorm:"foreignKey:TagID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"tag"`
}

func (ProblemTag) TableName() string {
	return "tbl_problem_tag"
}
