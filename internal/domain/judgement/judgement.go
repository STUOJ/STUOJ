package judgement

//go:generate go run ../../../dev/gen/query_gen.go judgement
//go:generate go run ../../../dev/gen/builder.go judgement

import (
	"STUOJ/internal/domain/judgement/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
)

// Judgement 表示判题记录领域对象
// 封装判题记录的核心业务逻辑和验证规则
type Judgement struct {
	Id            valueobject.Id
	SubmissionId  valueobject.SubmissionId
	TestcaseId    valueobject.TestcaseId
	Time          valueobject.Time
	Memory        valueobject.Memory
	Stdout        valueobject.Stdout
	Stderr        valueobject.Stderr
	CompileOutput valueobject.CompileOutput
	Message       valueobject.Message
	Status        valueobject.Status
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
