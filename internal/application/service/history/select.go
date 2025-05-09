package history

import (
	"STUOJ/internal/application/dto"
	"STUOJ/internal/application/dto/request"
	"STUOJ/internal/application/dto/response"
	"STUOJ/internal/domain/history"
	"STUOJ/internal/domain/user"
	querycontext2 "STUOJ/internal/infrastructure/persistence/repository/querycontext"
	query "STUOJ/internal/infrastructure/persistence/repository/queryfield"
)

type HistoryPage struct {
	Historys []response.HistoryListItemData `json:"historys"`
	dto.Page
}

func SelectById(id int64, reqUser request.ReqUser) (response.HistoryData, error) {

	var res response.HistoryData
	// 创建查询选项
	historyQuery := querycontext2.HistoryQueryContext{}
	historyQuery.Id.Add(id)
	historyQuery.Field = *query.HistoryAllField

	historyDomain, _, err := history.Query.SelectOne(historyQuery)
	if err != nil {
		return res, err
	}
	res = response.Domain2HistoryData(historyDomain)

	userQuery := querycontext2.UserQueryContext{}
	userQuery.Id.Add(historyDomain.UserId.Value())
	userQuery.Field = *query.UserSimpleField
	userDomain, _, err := user.Query.SelectOne(userQuery)
	if err == nil {
		res.User = response.Domain2UserSimpleData(userDomain)
	}
	return res, nil
}

// Select 查询所有历史记录
func Select(params request.QueryHistoryParams, reqUser request.ReqUser) (HistoryPage, error) {
	var res HistoryPage
	historyQueryContext := params2Query(params)
	historyQueryContext.Field = *query.HistorySimpleField

	historyDomains, _, err := history.Query.Select(historyQueryContext)
	if err != nil {
		return HistoryPage{}, err
	}

	// 构建历史记录列表
	res.Historys = make([]response.HistoryListItemData, len(historyDomains))
	for i, v := range historyDomains {
		res.Historys[i] = response.Domain2HistoryListItem(v)

		// 获取用户信息
		userQuery := querycontext2.UserQueryContext{}
		userQuery.Id.Add(v.UserId.Value())
		userQuery.Field = *query.UserSimpleField
		userDomain, _, err := user.Query.SelectOne(userQuery)
		if err == nil {
			res.Historys[i].User = response.Domain2UserSimpleData(userDomain)
		}
	}

	// 构建分页信息
	res.Page = dto.Page{
		Page:  historyQueryContext.Page.Page,
		Size:  historyQueryContext.Page.PageSize,
		Total: int64(len(historyDomains)),
	}

	return res, nil
}
