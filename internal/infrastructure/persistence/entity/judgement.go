package entity

import "time"

// JudgeStatus 评测状态：0 Pending, 1 In Queue, 2 Processing, 3 AC, 4 WA, 5 TLE, 6 CE, 7 RE(SIGSEGV), 8 RE(SIGXFSZ), 9 RE(SIGFPE), 10 RE(SIGABRT), 11 RE(NZEC), 12 RE(Other), 13 IE, 14 EFE
type JudgeStatus uint64

const (
	JudgePD        JudgeStatus = 0
	JudgeIQ        JudgeStatus = 1
	JudgePR        JudgeStatus = 2
	JudgeAC        JudgeStatus = 3
	JudgeWA        JudgeStatus = 4
	JudgeTLE       JudgeStatus = 5
	JudgeCE        JudgeStatus = 6
	JudgeRESIGSEGV JudgeStatus = 7
	JudgeRESIGXFSZ JudgeStatus = 8
	JudgeRESIGFPE  JudgeStatus = 9
	JudgeRESIGABRT JudgeStatus = 10
	JudgeRENZEC    JudgeStatus = 11
	JudgeRE        JudgeStatus = 12
	JudgeIE        JudgeStatus = 13
	JudgeEFE       JudgeStatus = 14
)

func (s JudgeStatus) String() string {
	switch s {
	case JudgePD:
		return "Pending"
	case JudgeIQ:
		return "In Queue"
	case JudgePR:
		return "Processing"
	case JudgeAC:
		return "Accepted"
	case JudgeWA:
		return "Wrong Answer"
	case JudgeTLE:
		return "Time Limit Exceeded"
	case JudgeCE:
		return "Compilation Error"
	case JudgeRESIGSEGV:
		return "Runtime Error (SIGSEGV)"
	case JudgeRESIGXFSZ:
		return "Runtime Error (SIGXFSZ)"
	case JudgeRESIGFPE:
		return "Runtime Error (SIGFPE)"
	case JudgeRESIGABRT:
		return "Runtime Error (SIGABRT)"
	case JudgeRENZEC:
		return "Runtime Error (NZEC)"
	case JudgeRE:
		return "Runtime Error"
	case JudgeIE:
		return "Internal Error"
	case JudgeEFE:
		return "Exec Format Error"
	default:
		return "Unknown"
	}
}

// Judgement 评测点结果
//
//go:generate go run ../../../../dev/gen/dao_store.go -struct=Judgement
//go:generate go run ../../../../dev/gen/field_select.go -struct=Judgement
type Judgement struct {
	Id            uint64      `gorm:"primaryKey;autoIncrement;comment:评测点Id"`
	SubmissionId  uint64      `gorm:"not null;default:0;comment:提交记录Id"`
	TestcaseId    uint64      `gorm:"not null;default:0;comment:评测点Id"`
	Time          float64     `gorm:"not null;default:0;comment:运行耗时（s）"`
	Memory        uint64      `gorm:"not null;default:0;comment:内存（kb）"`
	Stdout        string      `gorm:"type:longtext;not null;comment:标准输出"`
	Stderr        string      `gorm:"type:longtext;not null;comment:标准错误输出"`
	CompileOutput string      `gorm:"type:longtext;not null;comment:编译输出"`
	Message       string      `gorm:"type:longtext;not null;comment:信息"`
	CreateTime    time.Time   `gorm:"autoCreateTime;comment:创建时间"`
	UpdateTime    time.Time   `gorm:"autoUpdateTime;comment:更新时间"`
	Status        JudgeStatus `gorm:"not null;default:1;comment:状态"`
	Submission    Submission  `gorm:"foreignKey:SubmissionId;references:Id;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE"`
}

func (Judgement) TableName() string {
	return "tbl_judgement"
}
