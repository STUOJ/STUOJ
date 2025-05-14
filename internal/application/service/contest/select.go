package contest

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/contest"
	"STUOJ/internal/domain/user"
	"STUOJ/internal/infrastructure/persistence/repository/querycontext"
	"STUOJ/internal/infrastructure/persistence/repository/queryfield"
)

type ContestPage struct {
	Contests []response.ContestListItemData `json:"contests"`
	dto.Page
}

func Select(req request.QueryContestParams, reqUser request.ReqUser) (ContestPage, error) {
	var res ContestPage
	contestQuery := params2Query(req)
	contestQuery.Field = *queryfield.ContestSimpleField
	contestDomain, contestMap, err := contest.Query.Select(contestQuery, contest.QueryUserId())
	if err != nil {
		return res, err
	}
	userIds := make([]int64, len(contestDomain))
	for i, c := range contestMap {
		userId := c["contest_user_id"].(int64)
		userIds[i] = userId
	}
	userQuery := querycontext.UserQueryContext{}
	userQuery.Field = *queryfield.UserSimpleField
	userQuery.Id.Set(userIds)
	users, _, err := user.Query.SelectByIds(userQuery)
	if err != nil {
		return res, err
	}
	for _, c := range contestDomain {
		var resContest response.ContestListItemData
		resContest = domain2listItemResponse(c)
		resContest.User = response.Domain2UserSimpleData(users[c.UserId.Value()])
		res.Contests = append(res.Contests, resContest)
	}
	res.Page.Page = contestQuery.Page.Page
	res.Page.Size = contestQuery.Page.PageSize
	total, _ := GetStatistics(req)
	res.Page.Total = total
	return res, nil
}
