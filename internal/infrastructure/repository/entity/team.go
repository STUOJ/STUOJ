package entity

// TeamStatus 团队状态: 1 禁止, 2 无效, 3 有效
//
//go:generate go run ../../../utils/gen/enum_valid.go TeamStatus
type TeamStatus uint8

const (
	TeamBanned   TeamStatus = 1
	TeamDisabled TeamStatus = 2
	TeamEnabled  TeamStatus = 3
)

func (s TeamStatus) String() string {
	switch s {
	case TeamBanned:
		return "封禁"
	case TeamDisabled:
		return "无效"
	case TeamEnabled:
		return "有效"
	default:
		return "未知"
	}
}

// Team 团队
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=Team
//go:generate go run ../../../utils/gen/field_select.go -struct=Team
type Team struct {
	Id          uint64     `gorm:"primaryKey;autoIncrement;comment:团队Id"`
	UserId      uint64     `gorm:"not null;default:0;comment:用户Id"`
	ContestId   uint64     `gorm:"not null;default:0;comment:比赛Id"`
	Name        string     `gorm:"type:text;not null;comment:队名"`
	Description string     `gorm:"type:longtext;not null;comment:简介"`
	Status      TeamStatus `gorm:"not null;default:1;comment:状态"`
	User        User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Contest     Contest    `gorm:"foreignKey:ContestId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Team) TableName() string {
	return "tbl_team"
}

// TeamUser 团队用户关联
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=TeamUser
//go:generate go run ../../../utils/gen/field_select.go -struct=TeamUser
type TeamUser struct {
	TeamId uint64 `gorm:"primaryKey;not null;default:0;comment:团队Id"`
	UserId uint64 `gorm:"primaryKey;not null;default:0;comment:用户Id"`
	Team   Team   `gorm:"foreignKey:TeamId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   User   `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TeamUser) TableName() string {
	return "tbl_team_user"
}

// TeamSubmission 团队提交关联
//
//go:generate go run ../../../utils/gen/dao_store.go -struct=TeamSubmission
//go:generate go run ../../../utils/gen/field_select.go -struct=TeamSubmission
type TeamSubmission struct {
	TeamId       uint64     `gorm:"primaryKey;not null;default:0;comment:团队Id"`
	SubmissionId uint64     `gorm:"primaryKey;not null;default:0;comment:提交Id"`
	Team         Team       `gorm:"foreignKey:TeamId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Submission   Submission `gorm:"foreignKey:SubmissionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (TeamSubmission) TableName() string {
	return "tbl_team_submission"
}
