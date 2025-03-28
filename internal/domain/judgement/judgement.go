package judgement

import (
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type Judgement struct {
	Id            uint64
	SubmissionId  uint64
	TestcaseId    uint64
	Time          float64
	Memory        uint64
	Stdout        string
	Stderr        string
	CompileOutput string
	Message       string
	Status        entity.JudgeStatus
	CreateTime    time.Time
	UpdateTime    time.Time
}

func (j *Judgement) verify() error {
	return nil
}

func (j *Judgement) toEntity() entity.Judgement {
	return entity.Judgement{
		Id:            j.Id,
		SubmissionId:  j.SubmissionId,
		TestcaseId:    j.TestcaseId,
		Time:          j.Time,
		Memory:        j.Memory,
		Stdout:        j.Stdout,
		Stderr:        j.Stderr,
		CompileOutput: j.CompileOutput,
		Message:       j.Message,
		Status:        j.Status,
	}
}

func (j *Judgement) fromEntity(judgement entity.Judgement) *Judgement {
	j.Id = judgement.Id
	j.SubmissionId = judgement.SubmissionId
	j.TestcaseId = judgement.TestcaseId
	j.Time = judgement.Time
	j.Memory = judgement.Memory
	j.Stdout = judgement.Stdout
	j.Stderr = judgement.Stderr
	j.CompileOutput = judgement.CompileOutput
	j.Message = judgement.Message
	j.Status = judgement.Status
	return j
}

func (j *Judgement) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.JudgementId, option.OpEqual, j.Id)
	return options
}

func (j *Judgement) Create() (uint64, error) {
	j.CreateTime = time.Now()
	j.UpdateTime = time.Now()
	if err := j.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	judgement, err := dao.JudgementStore.Insert(j.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return judgement.Id, nil
}

func (j *Judgement) Update() error {
	var err error
	options := j.toOption()
	_, err = dao.JudgementStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	j.UpdateTime = time.Now()
	if err := j.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.JudgementStore.Updates(j.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (j *Judgement) Delete() error {
	options := j.toOption()
	_, err := dao.JudgementStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.JudgementStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

type Option func(*Judgement)

func NewJudgement(option ...Option) *Judgement {
	j := &Judgement{}
	for _, opt := range option {
		opt(j)
	}
	return j
}

func WithId(id uint64) Option {
	return func(j *Judgement) {
		j.Id = id
	}
}

func WithSubmissionId(submissionId uint64) Option {
	return func(j *Judgement) {
		j.SubmissionId = submissionId
	}
}

func WithTestcaseId(testcaseId uint64) Option {
	return func(j *Judgement) {
		j.TestcaseId = testcaseId
	}
}

func WithTime(time float64) Option {
	return func(j *Judgement) {
		j.Time = time
	}
}

func WithMemory(memory uint64) Option {
	return func(j *Judgement) {
		j.Memory = memory
	}
}

func WithStdout(stdout string) Option {
	return func(j *Judgement) {
		j.Stdout = stdout
	}
}

func WithStderr(stderr string) Option {
	return func(j *Judgement) {
		j.Stderr = stderr
	}
}

func WithCompileOutput(compileOutput string) Option {
	return func(j *Judgement) {
		j.CompileOutput = compileOutput
	}
}

func WithMessage(message string) Option {
	return func(j *Judgement) {
		j.Message = message
	}
}

func WithStatus(status entity.JudgeStatus) Option {
	return func(j *Judgement) {
		j.Status = status
	}
}
