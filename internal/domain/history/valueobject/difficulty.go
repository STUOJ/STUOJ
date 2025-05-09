package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/internal/infrastructure/persistence/entity"
	"STUOJ/pkg/errors"
)

type Difficulty struct {
	shared.Valueobject[entity.Difficulty]
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
