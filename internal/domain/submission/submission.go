package submission

//go:generate go run ../../../dev/gen/query_gen.go submission
//go:generate go run ../../../dev/gen/builder.go submission

import (
	"STUOJ/internal/infrastructure/repository/dao"
	entity "STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"time"

	"STUOJ/internal/domain/submission/valueobject"
)

type Submission struct {
	Id         int64
	UserId     int64
	ProblemId  int64
	Status     entity.JudgeStatus
	Score      int64
	Memory     int64
	Time       float64
	Length     int64
	LanguageId int64
	SourceCode valueobject.SourceCode
	CreateTime time.Time
	UpdateTime time.Time
}

func (s *Submission) verify() error {
	if err := s.SourceCode.Verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	if !entity.JudgeStatus(s.Status).IsValid() {
		return errors.ErrValidation.WithMessage("状态码无效")
	}
	return nil
}

func (s *Submission) toEntity() entity.Submission {
	return entity.Submission{
		Id:         uint64(s.Id),
		UserId:     uint64(s.UserId),
		ProblemId:  uint64(s.ProblemId),
		Status:     s.Status,
		Score:      uint8(s.Score),
		Memory:     uint64(s.Memory),
		Time:       s.Time,
		Length:     uint32(s.Length),
		LanguageId: uint64(s.LanguageId),
		SourceCode: s.SourceCode.String(),
		CreateTime: s.CreateTime,
		UpdateTime: s.UpdateTime,
	}
}

func (s *Submission) fromEntity(submission entity.Submission) *Submission {
	s.Id = int64(submission.Id)
	s.UserId = int64(submission.UserId)
	s.ProblemId = int64(submission.ProblemId)
	s.Status = submission.Status
	s.Score = int64(submission.Score)
	s.Memory = int64(submission.Memory)
	s.Time = submission.Time
	s.Length = int64(submission.Length)
	s.LanguageId = int64(submission.LanguageId)
	s.SourceCode = valueobject.NewSourceCode(submission.SourceCode)
	s.CreateTime = submission.CreateTime
	s.UpdateTime = submission.UpdateTime
	return s
}

func (s *Submission) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SubmissionId, option.OpEqual, s.Id)
	return options
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
