package solution

//go:generate go run ../../../dev/gen/query_gen.go solution
//go:generate go run ../../../dev/gen/builder.go solution

import (
	"STUOJ/internal/domain/solution/valueobject"
)

type Solution struct {
	Id         valueobject.Id
	LanguageId valueobject.LanguageId
	ProblemId  valueobject.ProblemId
	SourceCode valueobject.SourceCode
}
