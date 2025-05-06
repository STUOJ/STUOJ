package testcase

//go:generate go run ../../../dev/gen/query_gen.go testcase
//go:generate go run ../../../dev/gen/builder.go testcase

import (
	"STUOJ/internal/domain/testcase/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/pkg/errors"
)

type Testcase struct {
	Id         valueobject.Id
	ProblemId  valueobject.ProblemId
	Serial     valueobject.Serial
	TestInput  valueobject.TestInput
	TestOutput valueobject.TestOutput
}

func (t *Testcase) Create() (int64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	testcase, err := dao.TestcaseStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return int64(testcase.Id), nil
}

func (t *Testcase) Update() error {
	var err error
	options := t.toOption()
	_, err = dao.TestcaseStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	if err := t.verify(); err != nil {
		return errors.ErrValidation.WithMessage(err.Error())
	}
	_, err = dao.TestcaseStore.Updates(t.toEntity(), options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}

func (t *Testcase) Delete() error {
	options := t.toOption()
	_, err := dao.TestcaseStore.SelectOne(options)
	if err != nil {
		return errors.ErrNotFound.WithMessage(err.Error())
	}
	err = dao.TestcaseStore.Delete(options)
	if err != nil {
		return errors.ErrInternalServer.WithMessage(err.Error())
	}
	return nil
}
