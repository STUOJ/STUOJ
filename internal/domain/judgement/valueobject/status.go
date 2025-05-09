package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
)

type Status struct {
	model.Valueobject[entity.JudgeStatus]
}

func NewStatus(value entity.JudgeStatus) Status {
	var s Status
	s.Set(value)
	return s
}
