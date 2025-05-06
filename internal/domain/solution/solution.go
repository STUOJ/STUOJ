package solution

//go:generate go run ../../../dev/gen/query_gen.go solution
//go:generate go run ../../../dev/gen/builder.go solution

import (
	"STUOJ/internal/domain/solution/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
)

type Solution struct {
	Id         valueobject.Id
	LanguageId valueobject.LanguageId
	ProblemId  valueobject.ProblemId
	SourceCode valueobject.SourceCode
}

func (s *Solution) Create() (int64, error) {
	if err := s.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	solution, err := dao.SolutionStore.Insert(s.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(solution.Id), nil
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
	return nil
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
	return nil
}
