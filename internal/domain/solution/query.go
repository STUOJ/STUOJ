package solution

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
	"STUOJ/internal/model/querymodel"
)

type _Query struct{}

var Query = new(_Query)

func (query *_Query) Select(model querymodel.SolutionQueryModel) ([]Solution, error) {
	queryOptions := model.GenerateOptions()
	entitySolutions, err := dao.SolutionStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var solutions []Solution
	for _, entitySolution := range entitySolutions {
		solution := NewSolution().fromEntity(entitySolution)
		solutions = append(solutions, *solution)
	}
	return solutions, nil
}

func (query *_Query) SelectById(id uint64) (Solution, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.SolutionId, option.OpEqual, id)
	entitySolution, err := dao.SolutionStore.SelectOne(queryOptions)
	if err != nil {
		return Solution{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewSolution().fromEntity(entitySolution), nil
}

func (query *_Query) SelectByProblemId(problemId uint64) ([]Solution, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.SolutionProblemId, option.OpEqual, problemId)
	entitySolutions, err := dao.SolutionStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var solutions []Solution
	for _, entitySolution := range entitySolutions {
		solution := NewSolution().fromEntity(entitySolution)
		solutions = append(solutions, *solution)
	}
	return solutions, nil
}

func (query *_Query) SelectByLanguageId(languageId uint64) ([]Solution, error) {
	queryOptions := option.NewQueryOptions()
	queryOptions.Filters.Add(field.SolutionLanguageId, option.OpEqual, languageId)
	entitySolutions, err := dao.SolutionStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var solutions []Solution
	for _, entitySolution := range entitySolutions {
		solution := NewSolution().fromEntity(entitySolution)
		solutions = append(solutions, *solution)
	}
	return solutions, nil
}

func (query *_Query) Count(model querymodel.SolutionQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.SolutionStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
