package history

//go:generate go run ../../../dev/gen/query_gen.go history
//go:generate go run ../../../dev/gen/builder.go history

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/history/valueobject"
)

type History struct {
	Id           int64
	UserId       valueobject.UserId
	ProblemId    valueobject.ProblemId
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   entity.Difficulty
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

func (h *History) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.HistoryId, option.OpEqual, h.Id)
	return options
}

func (h *History) Create() (int64, error) {
	h.CreateTime = time.Now()
	if err := h.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	history, err := dao.HistoryStore.Insert(h.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(history.Id), nil
}

func (h *History) Update() error {
	var err error
	options := h.toOption()
	_, err = dao.HistoryStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := h.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.HistoryStore.Updates(h.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (h *History) Delete() error {
	options := h.toOption()
	_, err := dao.HistoryStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.HistoryStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
