package contest

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/contest"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/pkg/utils"
	"time"
)

func domain2listItemResponse(c contest.Contest) response.ContestListItemData {
	return response.ContestListItemData{
		Id:         c.Id.Value(),
		Title:      c.Title.Value(),
		Format:     int64(c.Format.Value()),
		Status:     int64(c.Status.Value()),
		StartTime:  c.StartTime.Value().String(),
		EndTime:    c.EndTime.Value().String(),
		TeamSize:   int64(c.TeamSize.Value()),
		CreateTime: c.CreateTime.String(),
		UpdateTime: c.UpdateTime.String(),
	}
}

func domain2response(c contest.Contest) response.ContestData {
	return response.ContestData{
		Id:          c.Id.Value(),
		Title:       c.Title.Value(),
		Format:      int64(c.Format.Value()),
		Status:      int64(c.Status.Value()),
		StartTime:   c.StartTime.Value().String(),
		EndTime:     c.EndTime.Value().String(),
		TeamSize:    int64(c.TeamSize.Value()),
		Description: c.Description.Value(),
		CreateTime:  c.CreateTime.String(),
		UpdateTime:  c.UpdateTime.String(),
	}
}

func params2Query(params request.QueryContestParams) (query querycontext.ContestQueryContext) {
	if params.UserId != nil {
		query.UserId.Add(*params.UserId)
	}
	if params.Title != nil {
		query.Title.Set(*params.Title)
	}
	if params.Status != nil {
		status, err := dao.StringToContestStatusSlice(*params.Status)
		if err == nil {
			query.Status.Set(status)
		}
	}
	if params.Format != nil {
		format, err := dao.StringToContestFormatSlice(*params.Format)
		if err == nil {
			query.Format.Set(format)
		}
	}
	if params.Status != nil {
		status, err := dao.StringToContestStatusSlice(*params.Status)
		if err == nil {
			query.Status.Set(status)
		}
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
	if params.TeamSize != nil {
		teamsize, err := utils.StringToInt8Slice(*params.TeamSize)
		if err == nil {
			query.TeamSize.Set(teamsize)
		}
	}
	if params.BeginStart != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.BeginStart)
		if err == nil {
			query.BeginStart.Set(t)
		}
	}
	if params.BeginEnd != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.BeginEnd)
		if err == nil {
			query.BeginEnd.Set(t)
		}
	}
	if params.FinishStart != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.FinishStart)
		if err == nil {
			query.FinishStart.Set(t)
		}
	}
	if params.FinishEnd != nil {
		t, err := time.Parse("2006-01-02 15:04:05", *params.FinishEnd)
		if err == nil {
			query.FinishEnd.Set(t)
		}
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.OrderBy, *params.Order)
	}
	return
}
