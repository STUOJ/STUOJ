package testcase

//go:generate go run ../../../dev/gen/query_gen.go testcase
//go:generate go run ../../../dev/gen/domain.go testcase

import (
	"STUOJ/internal/domain/testcase/valueobject"
)

type Testcase struct {
	Id         valueobject.Id
	ProblemId  valueobject.ProblemId
	Serial     valueobject.Serial
	TestInput  valueobject.TestInput
	TestOutput valueobject.TestOutput
}
