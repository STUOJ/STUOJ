package submission

//go:generate go run ../../../utils/gen/dto_gen.go submission
//go:generate go run ../../../utils/gen/query_gen.go submission

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/submission/valueobject"
	"STUOJ/internal/errors"
)

type Submission struct {
	Id         int64
	UserId     int64
	ProblemId  int64
	Status     entity.JudgeStatus
	Score      uint8
	Memory     int64
	Time       float64
	Length     uint32
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
	return &errors.NoError
}

func (s *Submission) toEntity() entity.Submission {
	return entity.Submission{
		Id:         uint64(s.Id),
		UserId:     uint64(s.UserId),
		ProblemId:  uint64(s.ProblemId),
		Status:     s.Status,
		Score:      s.Score,
		Memory:     uint64(s.Memory),
		Time:       s.Time,
		Length:     s.Length,
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
	s.Score = submission.Score
	s.Memory = int64(submission.Memory)
	s.Time = submission.Time
	s.Length = submission.Length
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
	return int64(submission.Id), &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}

type Option func(*Submission)

func NewSubmission(option ...Option) *Submission {
	s := &Submission{
		Status: entity.JudgeIE,
	}
	for _, opt := range option {
		opt(s)
	}
	return s
}

func WithId(id int64) Option {
	return func(s *Submission) {
		s.Id = id
	}
}

func WithUserId(userId int64) Option {
	return func(s *Submission) {
		s.UserId = userId
	}
}

func WithProblemId(problemId int64) Option {
	return func(s *Submission) {
		s.ProblemId = problemId
	}
}

func WithStatus(status entity.JudgeStatus) Option {
	return func(s *Submission) {
		s.Status = status
	}
}

func WithScore(score uint8) Option {
	return func(s *Submission) {
		s.Score = score
	}
}

func WithMemory(memory int64) Option {
	return func(s *Submission) {
		s.Memory = memory
	}
}

func WithTime(time float64) Option {
	return func(s *Submission) {
		s.Time = time
	}
}

func WithLength(length uint32) Option {
	return func(s *Submission) {
		s.Length = length
	}
}

func WithLanguageId(languageId int64) Option {
	return func(s *Submission) {
		s.LanguageId = languageId
	}
}

func WithSourceCode(sourceCode string) Option {
	return func(s *Submission) {
		s.SourceCode = valueobject.NewSourceCode(sourceCode)
	}
}
