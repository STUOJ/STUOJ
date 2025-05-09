package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"STUOJ/pkg/errors"
)

type Difficulty struct {
	model.Valueobject[entity.Difficulty]
}

func NewDifficulty(value entity.Difficulty) Difficulty {
	var s Difficulty
	s.Set(value)
	return s
}

func (s Difficulty) Verify() error {
	if !s.Value().IsValid() {
		return errors.ErrStatus
	}
	return nil
}
