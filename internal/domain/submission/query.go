package submission

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

func (*_Query) Select(model querymodel.SubmissionQueryModel) ([]Submission, error) {
	queryOptions := model.GenerateOptions()
	submissions, err := dao.SubmissionStore.Select(queryOptions)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Submission
	for _, submission := range submissions {
		result = append(result, *NewSubmission().fromEntity(submission))
	}
	return result, &errors.NoError
}

func (*_Query) SelectById(id uint64) (Submission, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SubmissionId, option.OpEqual, id)
	submission, err := dao.SubmissionStore.SelectOne(options)
	if err != nil {
		return Submission{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewSubmission().fromEntity(submission), &errors.NoError
}

func (*_Query) SelectSimpleById(id uint64) (Submission, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SubmissionId, option.OpEqual, id)
	options.Field = query.SubmissionSimpleField
	submission, err := dao.SubmissionStore.SelectOne(options)
	if err != nil {
		return Submission{}, errors.ErrNotFound.WithMessage(err.Error())
	}
	return *NewSubmission().fromEntity(submission), &errors.NoError
}

func (*_Query) SelectByUserId(userId uint64) ([]Submission, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SubmissionUserId, option.OpEqual, userId)
	options.Field = query.SubmissionAllField
	submissions, err := dao.SubmissionStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Submission
	for _, submission := range submissions {
		result = append(result, *NewSubmission().fromEntity(submission))
	}
	return result, &errors.NoError
}

func (*_Query) SelectByProblemId(problemId uint64) ([]Submission, error) {
	options := option.NewQueryOptions()
	options.Filters.Add(field.SubmissionProblemId, option.OpEqual, problemId)
	options.Field = query.SubmissionAllField
	submissions, err := dao.SubmissionStore.Select(options)
	if err != nil {
		return nil, errors.ErrInternalServer.WithMessage(err.Error())
	}
	var result []Submission
	for _, submission := range submissions {
		result = append(result, *NewSubmission().fromEntity(submission))
	}
	return result, &errors.NoError
}

func (*_Query) Count(model querymodel.SubmissionQueryModel) (int64, error) {
	queryOptions := model.GenerateOptions()
	count, err := dao.SubmissionStore.Count(queryOptions)
	if err != nil {
		return 0, errors.ErrInternalServer.WithMessage(err.Error())
	}
	return count, &errors.NoError
}
