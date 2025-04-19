package judgement

//go:generate go run ../../../utils/gen/dto_gen.go judgement
//go:generate go run ../../../utils/gen/query_gen.go judgement

import (
	"fmt"
	"time"

	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type Judgement struct {
	Id            int64
	SubmissionId  int64
	TestcaseId    int64
	Time          float64
	Memory        int64
	Stdout        string
	Stderr        string
	CompileOutput string
	Message       string
	Status        entity.JudgeStatus
	CreateTime    time.Time
	UpdateTime    time.Time
}

func (j *Judgement) verify() error {
	if j.SubmissionId == 0 {
		return fmt.Errorf("SubmissionId不能为空")
	}
	if j.TestcaseId == 0 {
		return fmt.Errorf("TestcaseId不能为空")
	}
	return nil
}

func (j *Judgement) toEntity() entity.Judgement {
	return entity.Judgement{
		Id:            uint64(j.Id),
		SubmissionId:  uint64(j.SubmissionId),
		TestcaseId:    uint64(j.TestcaseId),
		Time:          j.Time,
		Memory:        uint64(j.Memory),
		Stdout:        j.Stdout,
		Stderr:        j.Stderr,
		CompileOutput: j.CompileOutput,
		Message:       j.Message,
		Status:        j.Status,
	}
}

func (j *Judgement) fromEntity(judgement entity.Judgement) *Judgement {
	j.Id = int64(judgement.Id)
	j.SubmissionId = int64(judgement.SubmissionId)
	j.TestcaseId = int64(judgement.TestcaseId)
	j.Time = judgement.Time
	j.Memory = int64(judgement.Memory)
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

func (j *Judgement) Create() (int64, error) {
	j.CreateTime = time.Now()
	j.UpdateTime = time.Now()
	if err := j.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	judgement, err := dao.JudgementStore.Insert(j.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(judgement.Id), &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}

type Option func(*Judgement)

func NewJudgement(option ...Option) *Judgement {
	j := &Judgement{}
	for _, opt := range option {
		opt(j)
	}
	return j
}

func WithId(id int64) Option {
	return func(j *Judgement) {
		j.Id = id
	}
}

func WithSubmissionId(submissionId int64) Option {
	return func(j *Judgement) {
		j.SubmissionId = submissionId
	}
}

func WithTestcaseId(testcaseId int64) Option {
	return func(j *Judgement) {
		j.TestcaseId = testcaseId
	}
}

func WithTime(time float64) Option {
	return func(j *Judgement) {
		j.Time = time
	}
}

func WithMemory(memory int64) Option {
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
