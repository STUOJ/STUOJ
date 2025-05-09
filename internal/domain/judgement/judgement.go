package judgement

//go:generate go run ../../../dev/gen/query_gen.go judgement
//go:generate go run ../../../dev/gen/domain.go judgement

import (
	"STUOJ/internal/domain/judgement/valueobject"
)

// Judgement 表示判题记录领域对象
// 封装判题记录的核心业务逻辑和验证规则
type Judgement struct {
	Id            valueobject.Id
	SubmissionId  valueobject.SubmissionId
	TestcaseId    valueobject.TestcaseId
	Time          valueobject.Time
	Memory        valueobject.Memory
	Stdout        valueobject.Stdout
	Stderr        valueobject.Stderr
	CompileOutput valueobject.CompileOutput
	Message       valueobject.Message
	Status        valueobject.Status
}
