package valueobject

import (
	"STUOJ/internal/domain/shared"
	"STUOJ/internal/infrastructure/persistence/entity"
	"errors"
)

type Format struct {
	shared.Valueobject[entity.ContestFormat]
}

func NewFormat(value entity.ContestFormat) Format {
	var format Format
	format.Set(value)
	return format
}

func (f Format) Verify() error {
	if !f.Value().IsValid() {
		return errors.New("比赛格式无效")
	}
	return nil
}
