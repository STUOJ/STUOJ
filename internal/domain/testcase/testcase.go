package testcase

//go:generate go run ../../../utils/gen/dto_gen.go testcase
//go:generate go run ../../../utils/gen/query_gen.go testcase

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/domain/testcase/valueobject"
	"STUOJ/internal/errors"
)

type Testcase struct {
	Id         uint64
	ProblemId  uint64
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
		Id:         t.Id,
		ProblemId:  t.ProblemId,
		Serial:     t.Serial,
		TestInput:  t.TestInput.String(),
		TestOutput: t.TestOutput.String(),
	}
}

func (t *Testcase) fromEntity(testcase entity.Testcase) *Testcase {
	t.Id = testcase.Id
	t.ProblemId = testcase.ProblemId
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

func (t *Testcase) Create() (uint64, error) {
	if err := t.verify(); err != nil {
		return 0, errors.ErrValidation.WithMessage(err.Error())
	}
	testcase, err := dao.TestcaseStore.Insert(t.toEntity())
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return testcase.Id, &errors.NoError
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
	return &errors.NoError
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
	return &errors.NoError
}

type Option func(*Testcase)

func NewTestcase(option ...Option) *Testcase {
	t := &Testcase{}
	for _, opt := range option {
		opt(t)
	}
	return t
}

func WithId(id uint64) Option {
	return func(t *Testcase) {
		t.Id = id
	}
}

func WithProblemId(problemId uint64) Option {
	return func(t *Testcase) {
		t.ProblemId = problemId
	}
}

func WithSerial(serial uint16) Option {
	return func(t *Testcase) {
		t.Serial = serial
	}
}

func WithTestInput(testInput string) Option {
	return func(t *Testcase) {
		t.TestInput = valueobject.NewTestInput(testInput)
	}
}

func WithTestOutput(testOutput string) Option {
	return func(t *Testcase) {
		t.TestOutput = valueobject.NewTestOutput(testOutput)
	}
}
