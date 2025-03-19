package entity

// TeamStatus 团队状态: 1 禁止, 2 无效, 3 有效
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
type Team struct {
	Id          uint64     `gorm:"primaryKey;autoIncrement;comment:团队ID" json:"id"`
	UserId      uint64     `gorm:"not null;default:0;comment:用户ID" json:"-"`
	ContestId   uint64     `gorm:"not null;default:0;comment:比赛ID" json:"-"`
	Name        string     `gorm:"type:text;not null;comment:队名" json:"name"`
	Description string     `gorm:"type:longtext;not null;comment:简介" json:"description"`
	Status      TeamStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	User        User       `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Contest     Contest    `gorm:"foreignKey:ContestId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"contest"`
	UserIds     []uint64   `gorm:"-" json:"user_ids"`
	Users       []User     `gorm:"-" json:"users"`
}

func (Team) TableName() string {
	return "tbl_team"
}

// TeamUser 团队用户关联
type TeamUser struct {
	TeamId uint64 `gorm:"primaryKey;not null;default:0;comment:团队ID" json:"team_id"`
	UserId uint64 `gorm:"primaryKey;not null;default:0;comment:用户ID" json:"user_id"`
	Team   Team   `gorm:"foreignKey:TeamId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`
	User   User   `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (TeamUser) TableName() string {
	return "tbl_team_user"
}

// TeamSubmission 团队提交关联
type TeamSubmission struct {
	TeamId       uint64     `gorm:"primaryKey;not null;default:0;comment:团队ID" json:"team_id"`
	SubmissionId uint64     `gorm:"primaryKey;not null;default:0;comment:提交ID" json:"submission_id"`
	Team         Team       `gorm:"foreignKey:TeamId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`
	Submission   Submission `gorm:"foreignKey:SubmissionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"submission"`
}

func (TeamSubmission) TableName() string {
	return "tbl_team_submission"
}
