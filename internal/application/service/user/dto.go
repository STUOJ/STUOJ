package user

import (
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/repository/querycontext"
	"STUOJ/internal/model/option"
	"STUOJ/pkg/utils"
	"time"
)

func domain2Resp(dm user.User) (resp response.UserData) {
	resp = response.UserData{
		Id:         dm.Id,
		Username:   dm.Username.String(),
		Role:       uint8(dm.Role),
		Email:      dm.Email.String(),
		Avatar:     dm.Avatar.String(),
		Signature:  dm.Signature.String(),
		CreateTime: dm.CreateTime.String(),
		UpdateTime: dm.UpdateTime.String(),
	}
	return
}

func params2Query(params request.QueryUserParams) (query querycontext.UserQueryContext) {
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
	if params.Id != nil {
		ids, err := utils.StringToInt64Slice(*params.Id)
		if err != nil {
			query.Id.Set(ids)
		}
	}
	if params.Role != nil {
		ids, err := utils.StringToUint8Slice(*params.Role)
		if err != nil {
			query.Role.Set(ids)
		}
	}
	if params.Username != nil {
		query.Username.Set(*params.Username)
	}
	if params.Email != nil {
		query.Email.Set(*params.Email)
	}
	if params.Page != nil && params.Size != nil {
		query.Page = option.NewPagination(*params.Page, *params.Size)
	}
	if params.Order != nil && params.OrderBy != nil {
		query.Sort = option.NewSortQuery(*params.Order, *params.OrderBy)
	}
	return query
}
