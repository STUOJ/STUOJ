package record

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/judgement"
	"STUOJ/internal/domain/submission"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/utils"
	"time"
)

func domain2SubmissionData(dm submission.Submission) (resp response.SubmissionData) {
	resp = response.SubmissionData{
		Id:         dm.Id.Value(),
		Status:     int64(dm.Status.Value()),
		Score:      dm.Score.Value(),
		Memory:     dm.Memory.Value(),
		Time:       float64(dm.Time.Value()),
		Length:     dm.Length.Value(),
		LanguageId: dm.LanguageId.Value(),
		SourceCode: dm.SourceCode.Value(),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}

func domain2JudgementData(dm judgement.Judgement) (resp response.JudgementData) {
	resp = response.JudgementData{
		Id:            dm.Id.Value(),
		Memory:        dm.Memory.Value(),
		Message:       dm.Message.Value(),
		Status:        int64(dm.Status.Value()),
		Stderr:        dm.Stderr.Value(),
		Stdout:        dm.Stdout.Value(),
		SubmissionId:  dm.SubmissionId.Value(),
		CompileOutput: dm.CompileOutput.Value(),
		TestcaseId:    dm.TestcaseId.Value(),
		Time:          float64(dm.Time.Value()),
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
		ids, err := utils.StringToInt64Slice(*params.User)
		if err == nil {
			query.UserId.Set(ids)
		}
	}
	if params.Problem != nil {
		ids, err := utils.StringToInt64Slice(*params.Problem)
		if err == nil {
			query.ProblemId.Set(ids)
		}
	}
	if params.Status != nil {
		status, err := dao.StringToJudgeStatusSlice(*params.Status)
		if err == nil {
			query.Status.Set(status)
		}
	}
	if params.Language != nil {
		ids, err := utils.StringToInt64Slice(*params.Language)
		if err == nil {
			query.Language.Set(ids)
		}
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	return query
}
