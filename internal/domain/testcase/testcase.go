package testcase

//go:generate go run ../../../dev/gen/dto_gen.go testcase
//go:generate go run ../../../dev/gen/query_gen.go testcase
//go:generate go run ../../../dev/gen/builder.go testcase

import (
	"STUOJ/internal/domain/testcase/valueobject"
	"STUOJ/internal/infrastructure/repository/dao"
	"STUOJ/internal/infrastructure/repository/entity"
	"STUOJ/internal/infrastructure/repository/entity/field"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/errors"
)

type Testcase struct {
	Id         int64
	ProblemId  int64
	Serial     uint16
	TestInput  valueobject.TestInput
	TestOutput valueobject.TestOutput
}

func (t *Testcase) verify() error {
	if err := t.TestInput.Verify(); err != nil {
		return err
	}
	if err := t.TestOutput.Verify(); err != nil {
		return err
	}
	return nil
}

func (t *Testcase) toEntity() entity.Testcase {
	return entity.Testcase{
		Id:         uint64(t.Id),
		ProblemId:  uint64(t.ProblemId),
		Serial:     t.Serial,
		TestInput:  t.TestInput.String(),
		TestOutput: t.TestOutput.String(),
	}
}

func (t *Testcase) fromEntity(testcase entity.Testcase) *Testcase {
	t.Id = int64(testcase.Id)
	t.ProblemId = int64(testcase.ProblemId)
	t.Serial = testcase.Serial
	t.TestInput = valueobject.NewTestInput(testcase.TestInput)
	t.TestOutput = valueobject.NewTestOutput(testcase.TestOutput)
	return t
}

func (t *Testcase) toOption() *option.QueryOptions {
	options := option.NewQueryOptions()
	options.Filters.Add(field.TestcaseId, option.OpEqual, t.Id)
	return options
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
