package solution

//go:generate go run ../../../utils/gen/dto_gen.go solution
//go:generate go run ../../../utils/gen/query_gen.go solution

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/domain/solution/valueobject"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/option"
)

type Solution struct {
	Id         int64
	LanguageId int64
	ProblemId  int64
	SourceCode valueobject.SourceCode
}

func (s *Solution) verify() error {
	if err := s.SourceCode.Verify(); err != nil {
		return err
	}
	return nil
}

func (s *Solution) toEntity() entity.Solution {
	return entity.Solution{
		Id:         uint64(s.Id),
		LanguageId: uint64(s.LanguageId),
		ProblemId:  uint64(s.ProblemId),
		SourceCode: s.SourceCode.String(),
	}
}

func (s *Solution) fromEntity(solution entity.Solution) *Solution {
	s.Id = int64(solution.Id)
	s.LanguageId = int64(solution.LanguageId)
	s.ProblemId = int64(solution.ProblemId)
	s.SourceCode = valueobject.NewSourceCode(solution.SourceCode)
	return s
}

func (s *Solution) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SolutionId, option.OpEqual, s.Id)
	return options
}

func (s *Solution) Create() (int64, error) {
	if err := s.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	solution, err := dao.SolutionStore.Insert(s.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(solution.Id), &errors.NoError
}

func (s *Solution) Update() error {
	var err error
	options := s.toOption()
	_, err = dao.SolutionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := s.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.SolutionStore.Updates(s.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

func (s *Solution) Delete() error {
	options := s.toOption()
	_, err := dao.SolutionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.SolutionStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return &errors.NoError
}

type Option func(*Solution)

func NewSolution(option ...Option) *Solution {
	s := &Solution{}
	for _, opt := range option {
		opt(s)
	}
	return s
}

func WithId(id int64) Option {
	return func(s *Solution) {
		s.Id = id
	}
}

func WithLanguageId(languageId int64) Option {
	return func(s *Solution) {
		s.LanguageId = languageId
	}
}

func WithProblemId(problemId int64) Option {
	return func(s *Solution) {
		s.ProblemId = problemId
	}
}

func WithSourceCode(sourceCode string) Option {
	return func(s *Solution) {
		s.SourceCode = valueobject.NewSourceCode(sourceCode)
	}
}
