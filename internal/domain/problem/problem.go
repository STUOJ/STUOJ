package problem

//go:generate go run ../../../utils/gen/dto_gen.go problem

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/problem/valueobject"
	"STUOJ/internal/errors"
)

type Problem struct {
	Id           uint64
	Title        valueobject.Title
	Source       valueobject.Source
	Difficulty   entity.Difficulty
	TimeLimit    float64
	MemoryLimit  uint64
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
		Id:           p.Id,
		Title:        p.Title.String(),
		Source:       p.Source.String(),
		Difficulty:   p.Difficulty,
		TimeLimit:    p.TimeLimit,
		MemoryLimit:  p.MemoryLimit,
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
	p.Id = problem.Id
	p.Title = valueobject.NewTitle(problem.Title)
	p.Source = valueobject.NewSource(problem.Source)
	p.Difficulty = problem.Difficulty
	p.TimeLimit = problem.TimeLimit
	p.MemoryLimit = problem.MemoryLimit
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

func (p *Problem) Create() (uint64, error) {
	p.CreateTime = time.Now()
	p.UpdateTime = time.Now()
	if err := p.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	problem, err := dao.ProblemStore.Insert(p.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return problem.Id, &errors.NoError
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

func (p *Problem) UpdateTags(tagIds []uint64) error {
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
			ProblemId: p.Id,
			TagId:     id,
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

type Option func(*Problem)

func NewProblem(option ...Option) *Problem {
	p := &Problem{
		Status: entity.ProblemEditing,
	}
	for _, opt := range option {
		opt(p)
	}
	return p
}

func WithId(id uint64) Option {
	return func(p *Problem) {
		p.Id = id
	}
}

func WithTitle(title string) Option {
	return func(p *Problem) {
		p.Title = valueobject.NewTitle(title)
	}
}

func WithSource(source string) Option {
	return func(p *Problem) {
		p.Source = valueobject.NewSource(source)
	}
}

func WithDifficulty(difficulty entity.Difficulty) Option {
	return func(p *Problem) {
		p.Difficulty = difficulty
	}
}

func WithTimeLimit(timeLimit float64) Option {
	return func(p *Problem) {
		p.TimeLimit = timeLimit
	}
}

func WithMemoryLimit(memoryLimit uint64) Option {
	return func(p *Problem) {
		p.MemoryLimit = memoryLimit
	}
}

func WithDescription(description string) Option {
	return func(p *Problem) {
		p.Description = valueobject.NewDescription(description)
	}
}

func WithInput(input string) Option {
	return func(p *Problem) {
		p.Input = valueobject.NewInput(input)
	}
}

func WithOutput(output string) Option {
	return func(p *Problem) {
		p.Output = valueobject.NewOutput(output)
	}
}

func WithSampleInput(sampleInput string) Option {
	return func(p *Problem) {
		p.SampleInput = valueobject.NewInput(sampleInput)
	}
}

func WithSampleOutput(sampleOutput string) Option {
	return func(p *Problem) {
		p.SampleOutput = valueobject.NewOutput(sampleOutput)
	}
}

func WithHint(hint string) Option {
	return func(p *Problem) {
		p.Hint = valueobject.NewDescription(hint)
	}
}

func WithStatus(status entity.ProblemStatus) Option {
	return func(p *Problem) {
		p.Status = status
	}
}
