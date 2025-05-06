package problem

//go:generate go run ../../../dev/gen/query_gen.go problem
//go:generate go run ../../../dev/gen/builder.go problem

import (
	dao2 "STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/problem/valueobject"
)

type Problem struct {
	Id           valueobject.Id
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
	Status       valueobject.Status
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (p *Problem) UpdateTags(tagIds []int64) error {
	var err error
	options := p.toOption()
	_, err = dao2.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.ProblemTagStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range tagIds {
		_, err = dao2.ProblemTagStore.Insert(entity.ProblemTag{
			ProblemId: uint64(p.Id.Value()),
			TagId:     uint64(id),
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
	}
	return nil
}
