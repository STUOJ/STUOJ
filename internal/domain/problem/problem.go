package problem

//go:generate go run ../../../utils/gen/dto_gen.go problem
//go:generate go run ../../../utils/gen/query_gen.go problem

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/domain/problem/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/option"
)

type Problem struct {
	Id           int64
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   entity.Difficulty
	TimeLimit    float64
	MemoryLimit  int64
	Description  valueobject.Description
	Input        valueobject.Input
	Output       valueobject.Output
	SampleInput  valueobject.Input
	SampleOutput valueobject.Output
	Hint         valueobject.Description
	Status       entity.ProblemStatus
	CreateTime   time.Time
	UpdateTime   time.Time
}

func (p *Problem) verify() error {
	if err := p.Title.Verify(); err != nil {
		return err
	}
	if err := p.Source.Verify(); err != nil {
		return err
	}
	if err := p.Description.Verify(); err != nil {
		return err
	}
	if err := p.Input.Verify(); err != nil {
		return err
	}
	if err := p.Output.Verify(); err != nil {
		return err
	}
	if err := p.SampleInput.Verify(); err != nil {
		return err
	}
	if err := p.SampleOutput.Verify(); err != nil {
		return err
	}
	if err := p.Hint.Verify(); err != nil {
		return err
	}
	return nil
}

func (p *Problem) toEntity() entity.Problem {
	return entity.Problem{
		Id:           uint64(p.Id),
		Title:        p.Title.String(),
		Source:       p.Source.String(),
		Difficulty:   p.Difficulty,
		TimeLimit:    p.TimeLimit,
		MemoryLimit:  uint64(p.MemoryLimit),
		Description:  p.Description.String(),
		Input:        p.Input.String(),
		Output:       p.Output.String(),
		SampleInput:  p.SampleInput.String(),
		SampleOutput: p.SampleOutput.String(),
		Hint:         p.Hint.String(),
		Status:       p.Status,
		CreateTime:   p.CreateTime,
		UpdateTime:   p.UpdateTime,
	}
}

func (p *Problem) fromEntity(problem entity.Problem) *Problem {
	p.Id = int64(problem.Id)
	p.Title = valueobject.NewTitle(problem.Title)
	p.Source = valueobject.NewSource(problem.Source)
	p.Difficulty = problem.Difficulty
	p.TimeLimit = problem.TimeLimit
	p.MemoryLimit = int64(problem.MemoryLimit)
	p.Description = valueobject.NewDescription(problem.Description)
	p.Input = valueobject.NewInput(problem.Input)
	p.Output = valueobject.NewOutput(problem.Output)
	p.SampleInput = valueobject.NewInput(problem.SampleInput)
	p.SampleOutput = valueobject.NewOutput(problem.SampleOutput)
	p.Hint = valueobject.NewDescription(problem.Hint)
	p.Status = problem.Status
	p.CreateTime = problem.CreateTime
	p.UpdateTime = problem.UpdateTime
	return p
}

func (p *Problem) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.ProblemId, option.OpEqual, p.Id)
	return options
}

func (p *Problem) Create() (int64, error) {
	p.CreateTime = time.Now()
	p.UpdateTime = time.Now()
	if err := p.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	problem, err := dao.ProblemStore.Insert(p.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(problem.Id), &errors.NoError
}

func (p *Problem) Update() error {
	var err error
	options := p.toOption()
	_, err = dao.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	p.UpdateTime = time.Now()
	if err := p.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.ProblemStore.Updates(p.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (p *Problem) Delete() error {
	options := p.toOption()
	_, err := dao.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.ProblemStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (p *Problem) UpdateTags(tagIds []int64) error {
	var err error
	options := p.toOption()
	_, err = dao.ProblemStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.ProblemTagStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	var errs []error
	for _, id := range tagIds {
		_, err = dao.ProblemTagStore.Insert(entity.ProblemTag{
			ProblemId: uint64(p.Id),
			TagId:     uint64(id),
		})
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.ErrInternalServer.WithErrors(errs)
	}
	return &errors.NoError
}
