package testcase

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) Select(model querymodel.TestcaseQueryModel) ([]Testcase, error) {
	queryOptions := model.GenerateOptions()
	entityTestcases, err := dao.TestcaseStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var testcases []Testcase
	for _, entityTestcase := range entityTestcases {
		testcase := NewTestcase().fromEntity(entityTestcase)
		testcases = append(testcases, *testcase)
	}
	return testcases, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Testcase, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.TestcaseId, option.OpEqual, id)
	queryOptions.Field = query.TestcaseAllField
	entityTestcase, err := dao.TestcaseStore.SelectOne(queryOptions)
	if err != nil {
		return Testcase{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewTestcase().fromEntity(entityTestcase), &errors.NoError
}

func (*_Query) SelectByProblemId(problemId uint64) ([]Testcase, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.TestcaseProblemId, option.OpEqual, problemId)
	queryOptions.Field = query.TestcaseAllField
	entityTestcases, err := dao.TestcaseStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var testcases []Testcase
	for _, entityTestcase := range entityTestcases {
		testcase := NewTestcase().fromEntity(entityTestcase)
		testcases = append(testcases, *testcase)
	}
	return testcases, &errors.NoError
}

func (*_Query) Count(model querymodel.TestcaseQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.TestcaseStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
