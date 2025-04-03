package problem

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

func (*_Query) Select(model querymodel.ProblemQueryModel) ([]Problem, error) {
	queryOptions := model.GenerateOptions()
	entityProblems, err := dao.ProblemStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var problems []Problem
	for _, entityProblem := range entityProblems {
		problem := NewProblem().fromEntity(entityProblem)
		problems = append(problems, *problem)
	}
	return problems, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Problem, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.ProblemId, option.OpEqual, id)
	queryOptions.Field = query.ProblemAllField
	entityProblem, err := dao.ProblemStore.SelectOne(queryOptions)
	if err != nil {
		return Problem{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewProblem().fromEntity(entityProblem), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Problem, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.ProblemId, option.OpEqual, id)
	queryOptions.Field = query.ProblemSimpleField
	entityProblem, err := dao.ProblemStore.SelectOne(queryOptions)
	if err != nil {
		return Problem{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewProblem().fromEntity(entityProblem), &errors.NoError
}

func (*_Query) Count(model querymodel.ProblemQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.ProblemStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
