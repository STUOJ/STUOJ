package submission

//go:generate go run ../../../dev/gen/query_gen.go submission
//go:generate go run ../../../dev/gen/builder.go submission

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/submission/valueobject"
)

type Submission struct {
	Id         valueobject.Id
	UserId     valueobject.UserId
	ProblemId  valueobject.ProblemId
	Status     valueobject.Status
	Score      valueobject.Score
	Memory     valueobject.Memory
	Time       valueobject.Time
	Length     valueobject.Length
	LanguageId valueobject.LanguageId
	SourceCode valueobject.SourceCode
	CreateTime time.Time
	UpdateTime time.Time
}

func (s *Submission) Create() (int64, error) {
	s.CreateTime = time.Now()
	s.UpdateTime = time.Now()
	if err := s.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	submission, err := dao.SubmissionStore.Insert(s.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(submission.Id), nil
}

func (s *Submission) Update() error {
	var err error
	options := s.toOption()
	_, err = dao.SubmissionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	s.UpdateTime = time.Now()
	if err := s.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.SubmissionStore.Updates(s.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (s *Submission) Delete() error {
	options := s.toOption()
	_, err := dao.SubmissionStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.SubmissionStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
