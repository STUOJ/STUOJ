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

func (p *Problem) Create() (int64, error) {
	p.CreateTime = time.Now()
	p.UpdateTime = time.Now()
	if err := p.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	problem, err := dao2.ProblemStore.Insert(p.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(problem.Id), nil
}

func (p *Problem) Update() error {
	var err error
	options := p.toOption()
	_, err = dao2.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	p.UpdateTime = time.Now()
	if err := p.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao2.ProblemStore.Updates(p.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (p *Problem) Delete() error {
	options := p.toOption()
	_, err := dao2.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao2.ProblemStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
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
