package problem

import (
	"STUOJ/internal/domain/problem/valueobject"
	"STUOJ/internal/infrastructure/repository/entity"
)

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

func WithId(id int64) Option {
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

func WithMemoryLimit(memoryLimit int64) Option {
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
