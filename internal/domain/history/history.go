package history

//go:generate go run ../../../dev/gen/query_gen.go history
//go:generate go run ../../../dev/gen/domain.go history

import (
	"STUOJ/internal/infrastructure/persistence/entity"
	"time"

	"STUOJ/internal/domain/history/valueobject"
)

type History struct {
	Id           valueobject.Id
	UserId       valueobject.UserId
	ProblemId    valueobject.ProblemId
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   valueobject.Difficulty
	TimeLimit    valueobject.TimeLimit
	MemoryLimit  valueobject.MemoryLimit
	Description  valueobject.Description
	Input        valueobject.Input
	Output       valueobject.Output
	SampleInput  valueobject.Input
	SampleOutput valueobject.Output
	Hint         valueobject.Description
	Operation    entity.Operation
	CreateTime   time.Time
}
