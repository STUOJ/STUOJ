package valueobject

import (
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/model"
	"errors"
)

type Format struct {
	model.Valueobject[entity.ContestFormat]
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
