package judgement

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

func (*_Query) SelectById(id uint64) (Judgement, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.JudgementId, option.OpEqual, id)
	options.Field = query.JudgementAllField
	judgement, err := dao.JudgementStore.SelectOne(options)
	if err != nil {
		return Judgement{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewJudgement().fromEntity(judgement), &errors.NoError
}

func (*_Query) SelectBySubmissionId(submissionId uint64) ([]Judgement, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.JudgementSubmissionId, option.OpEqual, submissionId)
	options.Field = query.JudgementAllField
	judgements, err := dao.JudgementStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Judgement
	for _, judgement := range judgements {
		result = append(result, *NewJudgement().fromEntity(judgement))
	}
	return result, &errors.NoError
}

func (*_Query) Select(model querymodel.JudgementQueryModel) ([]Judgement, error) {
	queryOptions := model.GenerateOptions()
	queryOptions.Field = query.JudgementAllField
	judgements, err := dao.JudgementStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Judgement
	for _, judgement := range judgements {
		result = append(result, *NewJudgement().fromEntity(judgement))
	}
	return result, &errors.NoError
}

func (*_Query) Count(model querymodel.JudgementQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.JudgementStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
