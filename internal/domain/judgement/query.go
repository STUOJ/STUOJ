package judgement

import (
	"STUOJ/internal/db/dao"
	"STUOJ/internal/db/entity/field"
	"STUOJ/internal/db/query/option"
	"STUOJ/internal/errors"
)

type _Query struct{}

var Query = new(_Query)

func (*_Query) SelectById(id uint64) (Judgement, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.JudgementId, option.OpEqual, id)
	judgement, err := dao.JudgementStore.SelectOne(options)
	if err != nil {
		return Judgement{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewJudgement().fromEntity(judgement), nil
}

func (*_Query) SelectBySubmissionId(submissionId uint64) ([]Judgement, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.JudgementSubmissionId, option.OpEqual, submissionId)
	judgements, err := dao.JudgementStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Judgement
	for _, judgement := range judgements {
		result = append(result, *NewJudgement().fromEntity(judgement))
	}
	return result, nil
}

func (*_Query) Select(options *option.QueryOptions) ([]Judgement, error) {
	judgements, err := dao.JudgementStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Judgement
	for _, judgement := range judgements {
		result = append(result, *NewJudgement().fromEntity(judgement))
	}
	return result, nil
}

func (*_Query) Count(options *option.QueryOptions) (int64, error) {
	count, err := dao.JudgementStore.Count(options)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, nil
}
