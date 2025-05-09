package submission

//go:generate go run ../../../dev/gen/query_gen.go submission
//go:generate go run ../../../dev/gen/domain.go submission

import (
	"time"

	"STUOJ/internal/domain/submission/valueobject"
)

type Submission struct {
	Id         valueobject.Id
	UserId     valueobject.UserId
	ProblemId  valueobject.ProblemId
	Status     valueobject.Status
	Score      valueobject.Score
	Memory     valueobject.Memory
	Time       valueobject.Time
	Length     valueobject.Length
	LanguageId valueobject.LanguageId
	SourceCode valueobject.SourceCode
	CreateTime time.Time
	UpdateTime time.Time
}
