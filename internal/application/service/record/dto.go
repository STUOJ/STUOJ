package record

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/judgement"
	"STUOJ/internal/domain/submission"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"time"
)

func domain2SubmissionData(dm submission.Submission) (resp response.SubmissionData) {
	resp = response.SubmissionData{
		Id:         dm.Id,
		Status:     int64(dm.Status),
		Score:      dm.Score,
		Memory:     dm.Memory,
		Time:       dm.Time,
		Length:     dm.Length,
		LanguageId: dm.LanguageId,
		SourceCode: dm.SourceCode.String(),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}

func domain2JudgementData(dm judgement.Judgement) (resp response.JudgementData) {
	resp = response.JudgementData{
		Id:           dm.Id,
		Memory:       dm.Memory,
		Message:      dm.Message,
		Status:       int64(dm.Status),
		Stderr:       dm.Stderr,
		Stdout:       dm.Stdout,
		SubmissionId: dm.SubmissionId,
		TestcaseId:   dm.TestcaseId,
		Time:         dm.Time,
	}
	return
}

func params2Query(params request.QuerySubmissionParams) (query querycontext.SubmissionQueryContext) {
	if params.EndTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.EndTime)
		if err == nil {
			query.EndTime.Set(t)
		}
	}
	if params.StartTime != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.StartTime)
		if err == nil {
			query.StartTime.Set(t)
		}
	}
	if params.User != nil {
		query.UserId.Add(*params.User)
	}
	if params.Problem != nil {
		query.ProblemId.Add(*params.Problem)
	}
	if params.Status != nil {
		query.Status.Add(*params.Status)
	}
	if params.Language != nil {
		query.Language.Add(*params.Language)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	return query
}
