package problem

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/problem"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	option "STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/utils"
	"time"
)

func params2Query(params request.QueryProblemParams) (query querycontext.ProblemQueryContext) {
	if params.Difficulty != nil {
		if difficulty, err := dao.StringToDifficultySlice(*params.Difficulty); err == nil {
			query.Difficulty.Add(difficulty...)
		}
	}
	if params.Status != nil {
		if status, err := dao.StringToProblemStatusSlice(*params.Status); err == nil {
			query.Status.Set(status)
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
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
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
	return response.Domain2ProblemData(p)
}

func domain2listItemResponse(p problem.Problem) response.ProblemListItemData {
	return response.ProblemListItemData{
		Id:         p.Id.Value(),
		Title:      p.Title.Value(),
		Source:     p.Source.Value(),
		Status:     int64(p.Status.Value()),
		Difficulty: int64(p.Difficulty.Value()),
		CreateTime: p.CreateTime.String(),
		UpdateTime: p.UpdateTime.String(),
	}
}
