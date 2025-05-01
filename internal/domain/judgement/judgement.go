package judgement

//go:generate go run ../../../dev/gen/dto_gen.go judgement
//go:generate go run ../../../dev/gen/query_gen.go judgement
//go:generate go run ../../../dev/gen/builder.go judgement

import (
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
	"fmt"
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
	if err := j.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	judgement, err := dao.JudgementStore.Insert(j.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(judgement.Id), nil
}

func (j *Judgement) Update() error {
	var err error
	options := j.toOption()
	_, err = dao.JudgementStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
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
