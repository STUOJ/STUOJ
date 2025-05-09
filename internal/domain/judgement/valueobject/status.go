package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/internal/infrastructure/persistence/entity"
)

type Status struct {
	shared.Valueobject[entity.JudgeStatus]
}

func NewStatus(value entity.JudgeStatus) Status {
	var s Status
	s.Set(value)
	return s
}
