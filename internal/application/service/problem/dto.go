package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
	"time"
)

func params2Query(params request.QueryProblemParams) (query querycontext.ProblemQueryContext) {
	if params.Difficulty != nil {
		if difficulty, err := utils.StringToUint8Slice(*params.Difficulty); err == nil {
			query.Difficulty.Add(difficulty...)
		}
	}
	if params.Status != nil {
		if status, err := utils.StringToUint8Slice(*params.Status); err == nil {
			query.Status.Add(status...)
		}
	}
	if params.Title != nil {
		query.Title.Set(*params.Title)
	}
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
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	if params.Tag != nil {
		if tag, err := utils.StringToInt64Slice(*params.Tag); err == nil {
			problem.WhereTag(tag)(&query)
		}
	}
	if params.User != nil {
		problem.WhereUser(*params.User)(&query)
	}
	return query
}

func domain2response(p problem.Problem) response.ProblemData {
	return response.ProblemData{
		Id:           p.Id,
		Title:        p.Title.String(),
		Description:  p.Description.String(),
		Input:        p.Input.String(),
		Output:       p.Output.String(),
		Hint:         p.Hint.String(),
		Source:       p.Source.String(),
		TimeLimit:    int64(p.TimeLimit),
		MemoryLimit:  int64(p.MemoryLimit),
		Difficulty:   int64(p.Difficulty),
		Status:       int64(p.Status),
		SampleInput:  p.SampleInput.String(),
		SampleOutput: p.SampleOutput.String(),
		CreateTime:   p.CreateTime.String(),
		UpdateTime:   p.UpdateTime.String(),
	}
}

func domain2listItemResponse(p problem.Problem) response.ProblemListItemData {
	return response.ProblemListItemData{
		Id:         p.Id,
		Title:      p.Title.String(),
		Source:     p.Source.String(),
		Status:     int64(p.Status),
		Difficulty: int64(p.Difficulty),
		CreateTime: p.CreateTime.String(),
		UpdateTime: p.UpdateTime.String(),
	}
}
