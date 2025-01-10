package entity

// JudgeStatus 评测状态：0 Pend, 1 In Queue, 2 Proc, 3 AC, 4 WA, 5 TLE, 6 CE, 7 RE(SIGSEGV), 8 RE(SIGXFSZ), 9 RE(SIGFPE), 10 RE(SIGABRT), 11 RE(NZEC), 12 RE(Other), 13 IE, 14 EFE
type JudgeStatus uint64

const (
	JudgeStatusPend      JudgeStatus = 0
	JudgeStatusIQ        JudgeStatus = 1
	JudgeStatusProc      JudgeStatus = 2
	JudgeStatusAC        JudgeStatus = 3
	JudgeStatusWA        JudgeStatus = 4
	JudgeStatusTLE       JudgeStatus = 5
	JudgeStatusCE        JudgeStatus = 6
	JudgeStatusRESIGSEGV JudgeStatus = 7
	JudgeStatusRESIGXFSZ JudgeStatus = 8
	JudgeStatusRESIGFPE  JudgeStatus = 9
	JudgeStatusRESIGABRT JudgeStatus = 10
	JudgeStatusRENZEC    JudgeStatus = 11
	JudgeStatusREOther   JudgeStatus = 12
	JudgeStatusIE        JudgeStatus = 13
	JudgeStatusEFE       JudgeStatus = 14
)

func (s JudgeStatus) String() string {
	switch s {
	case JudgeStatusPend:
		return "Pending"
	case JudgeStatusIQ:
		return "In Queue"
	case JudgeStatusProc:
		return "Processing"
	case JudgeStatusAC:
		return "Accepted"
	case JudgeStatusWA:
		return "Wrong Answer"
	case JudgeStatusTLE:
		return "Time Limit Exceeded"
	case JudgeStatusCE:
		return "Compilation Error"
	case JudgeStatusRESIGSEGV:
		return "Runtime Error (SIGSEGV)"
	case JudgeStatusRESIGXFSZ:
		return "Runtime Error (SIGXFSZ)"
	case JudgeStatusRESIGFPE:
		return "Runtime Error (SIGFPE)"
	case JudgeStatusRESIGABRT:
		return "Runtime Error (SIGABRT)"
	case JudgeStatusRENZEC:
		return "Runtime Error (NZEC)"
	case JudgeStatusREOther:
		return "Runtime Error (Other)"
	case JudgeStatusIE:
		return "Internal Error"
	case JudgeStatusEFE:
		return "Exec Format Error"
	default:
		return "Unknown"
	}
}

// Judgement 评测点结果
type Judgement struct {
	ID            uint64      `gorm:"primaryKey;autoIncrement;comment:评测点ID" json:"id"`
	SubmissionID  uint64      `gorm:"not null;default:0;comment:提交记录ID" json:"submission_id"`
	TestcaseID    uint64      `gorm:"not null;default:0;comment:评测点ID" json:"testcase_id"`
	Time          float64     `gorm:"not null;default:0;comment:运行耗时（s）" json:"time"`
	Memory        uint64      `gorm:"not null;default:0;comment:内存（kb）" json:"memory"`
	Stdout        string      `gorm:"type:longtext;not null;comment:标准输出" json:"stdout"`
	Stderr        string      `gorm:"type:longtext;not null;comment:标准错误输出" json:"stderr"`
	CompileOutput string      `gorm:"type:longtext;not null;comment:编译输出" json:"compile_output"`
	Message       string      `gorm:"type:longtext;not null;comment:信息" json:"message"`
	Status        JudgeStatus `gorm:"not null;default:1;comment:状态" json:"status"`
	Submission    Submission  `gorm:"foreignKey:SubmissionID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE" json:"submission"`
}

func (Judgement) TableName() string {
	return "tbl_judgement"
}
